package postgresRepository

import (
	"context"
	"database/sql"

	postgresModel "remez_story/boundary/repository/postgres/model"
	"remez_story/domain/entity/chapter"
	loggerInterface "remez_story/infrastructure/logger/interface"
)

type ChapterRepository struct {
	db           *sql.DB
	logger       loggerInterface.Logger
	errProcessor *errorProcessor
}

func NewChapterRepository(db *sql.DB, logger loggerInterface.Logger) *ChapterRepository {
	return &ChapterRepository{
		db:           db,
		logger:       logger,
		errProcessor: newErrorProcessor(logger),
	}
}

func (r *ChapterRepository) GetByID(ctx context.Context, id chapter.ChapterID) (*chapter.Chapter, error) {
	const query = `SELECT id, title, description, order_index FROM chapters WHERE id = $1`

	var model postgresModel.ChapterModel
	err := r.db.QueryRowContext(ctx, query, id.Int64()).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.OrderIndex,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrChapterNotFound(ctx, id.Int64())
		}
		return nil, r.errProcessor.LogAndReturnErrFindChapters(ctx, err)
	}

	return postgresModel.ChapterFromModel(&model)
}

func (r *ChapterRepository) GetByOrderIndex(ctx context.Context, orderIndex int) (*chapter.Chapter, error) {
	const query = `SELECT id, title, description, order_index FROM chapters WHERE order_index = $1`

	var model postgresModel.ChapterModel
	err := r.db.QueryRowContext(ctx, query, orderIndex).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.OrderIndex,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrChapterNotFound(ctx, int64(orderIndex))
		}
		return nil, r.errProcessor.LogAndReturnErrFindChapters(ctx, err)
	}

	return postgresModel.ChapterFromModel(&model)
}

func (r *ChapterRepository) FindAll(ctx context.Context) ([]*chapter.Chapter, error) {
	const query = `SELECT id, title, description, order_index FROM chapters ORDER BY order_index ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, r.errProcessor.LogAndReturnErrFindChapters(ctx, err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error(ctx, closeErr)
		}
	}()

	chapters := make([]*chapter.Chapter, 0)
	for rows.Next() {
		var model postgresModel.ChapterModel
		if err := rows.Scan(
			&model.ID,
			&model.Title,
			&model.Description,
			&model.OrderIndex,
		); err != nil {
			r.logger.Error(ctx, err)
			continue
		}

		ch, err := postgresModel.ChapterFromModel(&model)
		if err != nil {
			r.logger.Error(ctx, err)
			continue
		}
		chapters = append(chapters, ch)
	}

	return chapters, nil
}
