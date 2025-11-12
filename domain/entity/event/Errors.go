package event

import "remez_story/infrastructure/errors"

var (
	ErrEventIDRequired              = errors.NewError("EVENT-001", "Event ID is required")
	ErrEventCodeRequired            = errors.NewError("EVENT-002", "Event code is required")
	ErrInvalidEventCode             = errors.NewError("EVENT-003", "Invalid event code format")
	ErrEventCodeUnsupportedScanType = errors.NewError("EVENT-004", "EventCode: unsupported Scan type")
)
