package repositoryInterface

import (
	"context"
	"remez_story/domain/entity/event"
)

type EventRepository interface {
	GetByID(ctx context.Context, id event.EventID) (*event.Event, error)
	GetByCode(ctx context.Context, code event.EventCode) (*event.Event, error)
	GetByIDs(ctx context.Context, ids []event.EventID) ([]*event.Event, error)
	FindAll(ctx context.Context) ([]*event.Event, error)
}
