package id

import "remez_story/infrastructure/errors"

var (
	ErrEntityIDIsEmpty     = errors.NewError("ID-001", "Entity id is empty")
	ErrInvalidEntityID     = errors.NewError("ID-002", "Entity id must be > 0")
	ErrUnsupportedScanType = errors.NewError("ID-003", "EntityID: unsupported Scan type")

	CreateEntityIDErrorCode = errors.ErrorCode("ID-003")
)

func ErrCreateEntityID(invalidID string) error {
	errMsg := "Fail create entityID from string = \"" + invalidID + "\""
	return errors.NewError(CreateEntityIDErrorCode, errMsg)
}
