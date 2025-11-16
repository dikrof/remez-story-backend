package postgresRepository

import (
	"context"
	"database/sql"

	postgresModel "remez_story/boundary/repository/postgres/model"
	"remez_story/domain/entity/player"
	loggerInterface "remez_story/infrastructure/logger/interface"
)

type ProgressRepository struct {
	db           *sql.DB
	logger       loggerInterface.Logger
	errProcessor *errorProcessor
}

func NewProgressRepository(db *sql.DB, logger loggerInterface.Logger) *ProgressRepository {
	return &ProgressRepository{
		db:           db,
		logger:       logger,
		errProcessor: newErrorProcessor(logger),
	}
}

func (r *ProgressRepository) Save(ctx context.Context, progress *player.Progress) error {
	model, err := postgresModel.ProgressToModel(progress)
	if err != nil {
		return r.errProcessor.LogAndReturnErrSaveProgress(ctx, err)
	}

	const query = `
INSERT INTO progress (
    player_id,
    current_node_id,
    state,
    decisions,
    money,
    reputation,
    started_at,
    updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    to_timestamp($7::double precision / 1000000000),
    to_timestamp($8::double precision / 1000000000)
)`
	_, err = r.db.ExecContext(
		ctx,
		query,
		model.PlayerID,
		model.CurrentNodeID,
		model.StateJSON,
		model.DecisionsJSON,
		model.Money,
		model.ReputationJSON,
		model.StartedAt,
		model.UpdatedAt,
	)
	if err != nil {
		return r.errProcessor.LogAndReturnErrSaveProgress(ctx, err)
	}
	return nil
}

func (r *ProgressRepository) GetByPlayerID(ctx context.Context, playerID player.PlayerID) (*player.Progress, error) {
	const query = `
SELECT
    player_id,
    current_node_id,
    state,
    decisions,
    money,
    reputation,
    EXTRACT(EPOCH FROM started_at)::BIGINT * 1000000000 AS started_at,
    EXTRACT(EPOCH FROM updated_at)::BIGINT * 1000000000 AS updated_at
FROM progress
WHERE player_id = $1
`
	var model postgresModel.ProgressModel
	err := r.db.QueryRowContext(ctx, query, playerID.String()).Scan(
		&model.PlayerID,
		&model.CurrentNodeID,
		&model.StateJSON,
		&model.DecisionsJSON,
		&model.Money,
		&model.ReputationJSON,
		&model.StartedAt,
		&model.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrProgressNotFound(ctx, playerID.String())
		}
		return nil, r.errProcessor.LogAndReturnErrSaveProgress(ctx, err)
	}

	return postgresModel.ProgressFromModel(&model)
}

func (r *ProgressRepository) Update(ctx context.Context, progress *player.Progress) error {
	model, err := postgresModel.ProgressToModel(progress)
	if err != nil {
		return r.errProcessor.LogAndReturnErrUpdateProgress(ctx, progress.PlayerID.String(), err)
	}

	const query = `
UPDATE progress
SET
    current_node_id = $2,
    state           = $3,
    decisions       = $4,
    money           = $5,
    reputation      = $6,
    updated_at      = to_timestamp($7::double precision / 1000000000)
WHERE player_id = $1
`
	result, err := r.db.ExecContext(
		ctx,
		query,
		model.PlayerID,
		model.CurrentNodeID,
		model.StateJSON,
		model.DecisionsJSON,
		model.Money,
		model.ReputationJSON,
		model.UpdatedAt,
	)
	if err != nil {
		return r.errProcessor.LogAndReturnErrUpdateProgress(ctx, progress.PlayerID.String(), err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return r.errProcessor.LogAndReturnErrProgressNotFound(ctx, progress.PlayerID.String())
	}

	return nil
}

func (r *ProgressRepository) Exists(ctx context.Context, playerID player.PlayerID) (bool, error) {
	const query = `SELECT EXISTS(SELECT 1 FROM progress WHERE player_id = $1)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, playerID.String()).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
