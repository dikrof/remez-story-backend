package event

import (
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
	"remez_story/infrastructure/errors"
)

type EventBuilder struct {
	event  *Event
	errors *errors.Errors
}

func NewEventBuilder() *EventBuilder {
	return &EventBuilder{
		event:  &Event{},
		errors: errors.NewErrors(),
	}
}

func (b *EventBuilder) ID(id EventID) *EventBuilder {
	b.event.ID = id
	return b
}

func (b *EventBuilder) Code(code EventCode) *EventBuilder {
	b.event.Code = code
	return b
}

func (b *EventBuilder) Title(titleText string) *EventBuilder {
	t, err := title.NewTitle(titleText)
	if err != nil {
		b.errors.AddError(err)
		return b
	}
	b.event.Title = t
	return b
}

func (b *EventBuilder) TitleValue(t title.Title) *EventBuilder {
	b.event.Title = t
	return b
}

func (b *EventBuilder) Description(descText string) *EventBuilder {
	d, err := description.NewDescription(descText)
	if err != nil {
		b.errors.AddError(err)
		return b
	}
	b.event.Description = d
	return b
}

func (b *EventBuilder) DescriptionValue(d description.Description) *EventBuilder {
	b.event.Description = d
	return b
}

func (b *EventBuilder) Deprecated(deprecated bool) *EventBuilder {
	b.event.Deprecated = deprecated
	return b
}

func (b *EventBuilder) Build() (*Event, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.event, nil
}

func (b *EventBuilder) checkRequiredFields() {
	if b.event.ID.IsZero() {
		b.errors.AddError(ErrEventIDRequired)
	}

	if b.event.Code.IsZero() {
		b.errors.AddError(ErrEventCodeRequired)
	}
}
