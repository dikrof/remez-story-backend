package chapter

import "remez_story/common/domainPrimitive/primitive/id"

type ChapterID struct {
	id.EntityID
}

func NewChapterID(v int64) (ChapterID, error) {
	x, err := id.NewEntityID(v)
	return ChapterID{EntityID: x}, err
}

func ChapterIDFromString(s string) (ChapterID, error) {
	x, err := id.EntityIDFrom(s)
	return ChapterID{EntityID: x}, err
}
