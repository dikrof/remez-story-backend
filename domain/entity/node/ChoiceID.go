package node

import "remez_story/common/domainPrimitive/primitive/id"

type ChoiceID struct {
	id.EntityID
}

func NewChoiceID(v int64) (ChoiceID, error) {
	x, err := id.NewEntityID(v)
	return ChoiceID{EntityID: x}, err
}

func ChoiceIDFromString(s string) (ChoiceID, error) {
	x, err := id.EntityIDFrom(s)
	return ChoiceID{EntityID: x}, err
}
