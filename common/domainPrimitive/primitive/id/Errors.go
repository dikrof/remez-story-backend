package id

import (
	"fmt"
	"remez_story/infrastructure/errors"
)

var (
	ErrInvalidEntityID     = errors.NewError("ID-001", "Invalid entity ID: must be positive")
	ErrEntityIDIsEmpty     = errors.NewError("ID-002", "Entity ID string is empty")
	ErrUnsupportedScanType = errors.NewError("ID-003", "EntityID: unsupported Scan type")
)

func ErrCreateEntityID(value string) *errors.Error {
	return errors.NewError("ID-004", fmt.Sprintf("Failed to create EntityID from string: %q", value))
}
