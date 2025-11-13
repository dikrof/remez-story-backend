package description

import "remez_story/infrastructure/errors"

var (
	ErrDescriptionTooLong             = errors.NewError("DESC-001", "Description is too long (max 2000 chars)")
	ErrDescriptionUnsupportedScanType = errors.NewError("DESC-002", "Description: unsupported Scan type")
)
