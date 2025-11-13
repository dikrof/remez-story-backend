package node

import "remez_story/infrastructure/errors"

var (
	ErrNodeIDRequired                = errors.NewError("NODE-001", "Node ID is required")
	ErrInvalidNodeKind               = errors.NewError("NODE-002", "Invalid node kind")
	ErrChoicesRequired               = errors.NewError("NODE-004", "Choice node must have choices")
	ErrTextRequired                  = errors.NewError("NODE-005", "Node text is required")
	ErrChapterIDRequired             = errors.NewError("NODE-006", "Chapter ID is required")
	ErrSceneLabelTooLong             = errors.NewError("NODE-007", "SceneLabel too long (max 128 chars)")
	ErrChoiceIDRequired              = errors.NewError("NODE-008", "Choice ID is required")
	ErrChoiceTextRequired            = errors.NewError("NODE-009", "Choice text is required")
	ErrChoiceToNodeRequired          = errors.NewError("NODE-010", "Choice ToNodeID is required")
	ErrNodeKindUnsupportedScanType   = errors.NewError("NODE-011", "NodeKind: unsupported Scan type")
	ErrSceneLabelUnsupportedScanType = errors.NewError("NODE-012", "SceneLabel: unsupported Scan type")
	ErrChoiceToNodeIDRequired        = errors.NewError("NODE-013", "Choice ToNodeID is required")
	ErrChoiceTextIsRequired          = errors.NewError("NODE-014", "Choice text is required")
	ErrConditionalToNodeIDRequired   = errors.NewError("NODE-016", "Conditional ToNodeID is required")
	ErrInvalidNextID                 = errors.NewError("NODE-017", "This node kind cannot have NextID")
)
