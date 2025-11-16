package chapter

import (
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
)

type Chapter struct {
	ID          ChapterID
	Title       title.Title
	Description description.Description
	OrderIndex  int
}

func (c *Chapter) GetID() ChapterID {
	return c.ID
}

func (c *Chapter) GetTitle() title.Title {
	return c.Title
}

func (c *Chapter) GetDescription() description.Description {
	return c.Description
}

func (c *Chapter) GetOrderIndex() int {
	return c.OrderIndex
}
