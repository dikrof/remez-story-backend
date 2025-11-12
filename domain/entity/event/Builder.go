package event

import "remez_story/infrastructure/errors"

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

func (b *EventBuilder) Title(title string) *EventBuilder {
	b.event.Title = title
	return b
}

func (b *EventBuilder) Description(description string) *EventBuilder {
	b.event.Description = description
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
