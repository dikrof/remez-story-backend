package chapter

import "remez_story/infrastructure/errors"

var (
	ErrChapterIDRequired = errors.NewError("CHAPTER-001", "Chapter ID is required")
	ErrTitleRequired     = errors.NewError("CHAPTER-002", "Chapter title is required")
)
