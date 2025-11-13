package id

import (
	"fmt"
	"remez_story/infrastructure/errors"
)

var (
	ErrEntityIDIsEmpty     = errors.NewError("02e4e913-001", "Entity id is empty")
	ErrInvalidEntityID     = errors.NewError("02e4e913-002", "Entity id must be > 0")
	ErrUnsupportedScanType = errors.NewError("02e4e913-003", "EntityID: unsupported Scan type")

	CreateEntityIDErrorCode = errors.ErrorCode("02e4e913-004")
)

func ErrCreateEntityID(invalidID string) error {
	errMsg := fmt.Sprintf("Fail create entityID from string = %q", invalidID)
	return errors.NewError(CreateEntityIDErrorCode, errMsg)
}
