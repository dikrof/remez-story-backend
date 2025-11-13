package title

import "remez_story/infrastructure/errors"

var (
	ErrTitleIsEmpty             = errors.NewError("TITLE-001", "Title cannot be empty")
	ErrTitleTooLong             = errors.NewError("TITLE-002", "Title is too long (max 240 chars)")
	ErrTitleUnsupportedScanType = errors.NewError("TITLE-003", "Title: unsupported Scan type")
)
