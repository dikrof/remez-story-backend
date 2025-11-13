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
