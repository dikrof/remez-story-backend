package postgresRepository

import (
	"context"
	"database/sql"

	postgresModel "remez_story/boundary/repository/postgres/model"
	"remez_story/domain/entity/chapter"
	"remez_story/domain/entity/event"
	"remez_story/domain/entity/node"
	loggerInterface "remez_story/infrastructure/logger/interface"
)

type NodeRepository struct {
	db           *sql.DB
	logger       loggerInterface.Logger
	errProcessor *errorProcessor
}

func NewNodeRepository(db *sql.DB, logger loggerInterface.Logger) *NodeRepository {
	return &NodeRepository{
		db:           db,
		logger:       logger,
		errProcessor: newErrorProcessor(logger),
	}
}

func (r *NodeRepository) GetByID(ctx context.Context, id node.NodeID) (*node.Node, error) {
	const query = `SELECT id, chapter_id, scene_label, kind, speaker, text, next_id, choices, conditional
FROM nodes WHERE id = $1`

	var model postgresModel.NodeModel
	err := r.db.QueryRowContext(ctx, query, id.Int64()).Scan(
		&model.ID,
		&model.ChapterID,
		&model.SceneLabel,
		&model.Kind,
		&model.Speaker,
		&model.Text,
		&model.NextID,
		&model.ChoicesJSON,
		&model.ConditionalJSON,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrNodeNotFound(ctx, id.Int64())
		}
		return nil, r.errProcessor.LogAndReturnErrFindNodes(ctx, err)
	}

	return postgresModel.NodeFromModel(&model)
}

func (r *NodeRepository) GetByIDs(ctx context.Context, ids []node.NodeID) ([]*node.Node, error) {
	if len(ids) == 0 {
		return []*node.Node{}, nil
	}

	int64IDs := make([]int64, 0, len(ids))
	for _, id := range ids {
		int64IDs = append(int64IDs, id.Int64())
	}

	const query = `SELECT id, chapter_id, scene_label, kind, speaker, text, next_id, choices, conditional
FROM nodes WHERE id = ANY($1) ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query, int64IDs)
	if err != nil {
		return nil, r.errProcessor.LogAndReturnErrFindNodes(ctx, err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error(ctx, closeErr)
		}
	}()

	nodesRes := make([]*node.Node, 0, len(ids))
	for rows.Next() {
		var model postgresModel.NodeModel
		if err := rows.Scan(
			&model.ID,
			&model.ChapterID,
			&model.SceneLabel,
			&model.Kind,
			&model.Speaker,
			&model.Text,
			&model.NextID,
			&model.ChoicesJSON,
			&model.ConditionalJSON,
		); err != nil {
			r.logger.Error(ctx, err)
			continue
		}

		n, err := postgresModel.NodeFromModel(&model)
		if err != nil {
			r.logger.Error(ctx, err)
			continue
		}
		nodesRes = append(nodesRes, n)
	}

	return nodesRes, nil
}

func (r *NodeRepository) GetStartNodeForChapter(ctx context.Context, chapterID chapter.ChapterID) (*node.Node, error) {
	const query = `SELECT id, chapter_id, scene_label, kind, speaker, text, next_id, choices, conditional
FROM nodes WHERE chapter_id = $1 ORDER BY id ASC LIMIT 1`

	var model postgresModel.NodeModel
	err := r.db.QueryRowContext(ctx, query, chapterID.Int64()).Scan(
		&model.ID,
		&model.ChapterID,
		&model.SceneLabel,
		&model.Kind,
		&model.Speaker,
		&model.Text,
		&model.NextID,
		&model.ChoicesJSON,
		&model.ConditionalJSON,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrStartNodeNotFound(ctx, chapterID.Int64())
		}
		return nil, r.errProcessor.LogAndReturnErrFindNodes(ctx, err)
	}

	return postgresModel.NodeFromModel(&model)
}

func (r *NodeRepository) GetNextPossibleNodes(
	ctx context.Context,
	currentNodeID node.NodeID,
	playerState map[event.EventID]struct{},
) ([]*node.Node, error) {
	const query = `
WITH current_node AS (
    SELECT * FROM nodes WHERE id = $1
),
next_node_ids AS (
    SELECT unnest(get_possible_next_nodes(current_node.*)) AS node_id
    FROM current_node
)
SELECT n.id,
       n.chapter_id,
       n.scene_label,
       n.kind,
       n.speaker,
       n.text,
       n.next_id,
       n.choices,
       n.conditional
FROM nodes n
JOIN next_node_ids nni ON n.id = nni.node_id
`

	rows, err := r.db.QueryContext(ctx, query, currentNodeID.Int64())
	if err != nil {
		return nil, r.errProcessor.LogAndReturnErrFindNodes(ctx, err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error(ctx, closeErr)
		}
	}()

	nodesRes := make([]*node.Node, 0)
	for rows.Next() {
		var model postgresModel.NodeModel
		if err := rows.Scan(
			&model.ID,
			&model.ChapterID,
			&model.SceneLabel,
			&model.Kind,
			&model.Speaker,
			&model.Text,
			&model.NextID,
			&model.ChoicesJSON,
			&model.ConditionalJSON,
		); err != nil {
			r.logger.Error(ctx, err)
			continue
		}

		n, err := postgresModel.NodeFromModel(&model)
		if err != nil {
			r.logger.Error(ctx, err)
			continue
		}

		if r.shouldIncludeNode(n, playerState) {
			nodesRes = append(nodesRes, n)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, r.errProcessor.LogAndReturnErrFindNodes(ctx, err)
	}

	return nodesRes, nil
}

func (r *NodeRepository) shouldIncludeNode(n *node.Node, playerState map[event.EventID]struct{}) bool {
	if len(n.Conditional) == 0 {
		return true
	}

	for _, edge := range n.Conditional {
		allPresent := true
		for _, reqEventID := range edge.When.RequireAll {
			if _, exists := playerState[reqEventID]; !exists {
				allPresent = false
				break
			}
		}

		nonePresent := true
		for _, reqEventID := range edge.When.RequireNone {
			if _, exists := playerState[reqEventID]; exists {
				nonePresent = false
				break
			}
		}

		if allPresent && nonePresent {
			return true
		}
	}

	return false
}
