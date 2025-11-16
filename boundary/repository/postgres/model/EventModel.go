package postgresModel

import (
	"database/sql"
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
	"remez_story/domain/entity/event"
	"remez_story/infrastructure/errors"
)

type EventModel struct {
	ID          int64          `db:"id"`
	Code        string         `db:"code"`
	Title       string         `db:"title"`
	Description sql.NullString `db:"description"`
	Deprecated  bool           `db:"deprecated"`
}

func EventFromModel(model *EventModel) (*event.Event, error) {
	errs := errors.NewErrors()

	eventID, err := event.NewEventID(model.ID)
	if err != nil {
		errs.AddError(err)
	}

	eventCode, err := event.NewEventCode(model.Code)
	if err != nil {
		errs.AddError(err)
	}

	eventTitle, err := title.NewTitle(model.Title)
	if err != nil {
		errs.AddError(err)
	}

	if errs.IsPresent() {
		return nil, errs
	}

	builder := event.NewEventBuilder().
		ID(eventID).
		Code(eventCode).
		TitleValue(eventTitle).
		Deprecated(model.Deprecated)

	if model.Description.Valid {
		desc, err := description.NewDescription(model.Description.String)
		if err != nil {
			return nil, err
		}
		builder.DescriptionValue(desc)
	}

	return builder.Build()
}
