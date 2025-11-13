package event

import (
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
)

type Event struct {
	ID          EventID
	Code        EventCode
	Title       title.Title
	Description description.Description
	Deprecated  bool
}

func (e *Event) GetID() EventID {
	return e.ID
}

func (e *Event) GetCode() EventCode {
	return e.Code
}

func (e *Event) GetTitle() title.Title {
	return e.Title
}

func (e *Event) GetDescription() description.Description {
	return e.Description
}

func (e *Event) IsDeprecated() bool {
	return e.Deprecated
}

func (e *Event) Deprecate() {
	e.Deprecated = true
}

func (e *Event) Restore() {
	e.Deprecated = false
}
