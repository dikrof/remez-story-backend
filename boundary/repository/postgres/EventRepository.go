package postgresRepository

import (
	"context"
	"database/sql"

	postgresModel "remez_story/boundary/repository/postgres/model"
	"remez_story/domain/entity/event"
	loggerInterface "remez_story/infrastructure/logger/interface"
)

type EventRepository struct {
	db           *sql.DB
	logger       loggerInterface.Logger
	errProcessor *errorProcessor
}

func NewEventRepository(db *sql.DB, logger loggerInterface.Logger) *EventRepository {
	return &EventRepository{
		db:           db,
		logger:       logger,
		errProcessor: newErrorProcessor(logger),
	}
}

func (r *EventRepository) GetByID(ctx context.Context, id event.EventID) (*event.Event, error) {
	const query = `SELECT id, code, title, description, deprecated FROM events WHERE id = $1`

	var model postgresModel.EventModel
	err := r.db.QueryRowContext(ctx, query, id.Int64()).Scan(
		&model.ID,
		&model.Code,
		&model.Title,
		&model.Description,
		&model.Deprecated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrEventNotFound(ctx, id.Int64())
		}
		return nil, r.errProcessor.LogAndReturnErrFindEvents(ctx, err)
	}

	return postgresModel.EventFromModel(&model)
}

func (r *EventRepository) GetByCode(ctx context.Context, code event.EventCode) (*event.Event, error) {
	const query = `SELECT id, code, title, description, deprecated FROM events WHERE code = $1`

	var model postgresModel.EventModel
	err := r.db.QueryRowContext(ctx, query, code.String()).Scan(
		&model.ID,
		&model.Code,
		&model.Title,
		&model.Description,
		&model.Deprecated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.errProcessor.LogAndReturnErrEventNotFound(ctx, 0)
		}
		return nil, r.errProcessor.LogAndReturnErrFindEvents(ctx, err)
	}

	return postgresModel.EventFromModel(&model)
}

func (r *EventRepository) GetByIDs(ctx context.Context, ids []event.EventID) ([]*event.Event, error) {
	if len(ids) == 0 {
		return []*event.Event{}, nil
	}

	int64IDs := make([]int64, 0, len(ids))
	for _, id := range ids {
		int64IDs = append(int64IDs, id.Int64())
	}

	const query = `SELECT id, code, title, description, deprecated FROM events WHERE id = ANY($1)`

	rows, err := r.db.QueryContext(ctx, query, int64IDs)
	if err != nil {
		return nil, r.errProcessor.LogAndReturnErrFindEvents(ctx, err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error(ctx, closeErr)
		}
	}()

	events := make([]*event.Event, 0, len(ids))
	for rows.Next() {
		var model postgresModel.EventModel
		if err := rows.Scan(
			&model.ID,
			&model.Code,
			&model.Title,
			&model.Description,
			&model.Deprecated,
		); err != nil {
			r.logger.Error(ctx, err)
			continue
		}

		e, err := postgresModel.EventFromModel(&model)
		if err != nil {
			r.logger.Error(ctx, err)
			continue
		}
		events = append(events, e)
	}

	return events, nil
}

func (r *EventRepository) FindAll(ctx context.Context) ([]*event.Event, error) {
	const query = `SELECT id, code, title, description, deprecated FROM events ORDER BY id ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, r.errProcessor.LogAndReturnErrFindEvents(ctx, err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error(ctx, closeErr)
		}
	}()

	events := make([]*event.Event, 0)
	for rows.Next() {
		var model postgresModel.EventModel
		if err := rows.Scan(
			&model.ID,
			&model.Code,
			&model.Title,
			&model.Description,
			&model.Deprecated,
		); err != nil {
			r.logger.Error(ctx, err)
			continue
		}

		e, err := postgresModel.EventFromModel(&model)
		if err != nil {
			r.logger.Error(ctx, err)
			continue
		}
		events = append(events, e)
	}

	return events, nil
}
