package event

import "remez_story/common/domainPrimitive/primitive/id"

type EventID struct {
	id.EntityID
}

func NewEventID(v int64) (EventID, error) {
	x, err := id.NewEntityID(v)
	return EventID{EntityID: x}, err
}

func EventIDFromString(s string) (EventID, error) {
	x, err := id.EntityIDFrom(s)
	return EventID{EntityID: x}, err
}
