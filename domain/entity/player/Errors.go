package player

import "remez_story/infrastructure/errors"

var (
	ErrPlayerIDEmpty               = errors.NewError("PLAYER-001", "PlayerID must be non-empty")
	ErrPlayerIDTooLong             = errors.NewError("PLAYER-002", "PlayerID too long (max 128 chars)")
	ErrPlayerIDRequired            = errors.NewError("PLAYER-003", "Player ID is required")
	ErrPlayerIDUnsupportedScanType = errors.NewError("PLAYER-004", "PlayerID: unsupported Scan type")
)
