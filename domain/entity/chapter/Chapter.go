package chapter

import (
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
)

type Chapter struct {
	ID          ChapterID
	Title       title.Title
	Description description.Description
}
