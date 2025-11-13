package node

import "remez_story/infrastructure/errors"

var (
	ErrNodeIDRequired                = errors.NewError("NODE-001", "Node ID is required")
	ErrInvalidNodeKind               = errors.NewError("NODE-002", "Invalid node kind")
	ErrChoicesRequired               = errors.NewError("NODE-003", "Choice node must have choices")
	ErrTextRequired                  = errors.NewError("NODE-004", "Node text is required")
	ErrChapterIDRequired             = errors.NewError("NODE-005", "Chapter ID is required")
	ErrSceneLabelTooLong             = errors.NewError("NODE-006", "SceneLabel too long (max 128 chars)")
	ErrChoiceIDRequired              = errors.NewError("NODE-007", "Choice ID is required")
	ErrChoiceTextRequired            = errors.NewError("NODE-008", "Choice text is required")
	ErrChoiceToNodeRequired          = errors.NewError("NODE-009", "Choice ToNodeID is required")
	ErrNodeKindUnsupportedScanType   = errors.NewError("NODE-010", "NodeKind: unsupported Scan type")
	ErrSceneLabelUnsupportedScanType = errors.NewError("NODE-011", "SceneLabel: unsupported Scan type")
	ErrChoiceToNodeIDRequired        = errors.NewError("NODE-012", "Choice ToNodeID is required")
	ErrChoiceTextIsRequired          = errors.NewError("NODE-013", "Choice text is required")
	ErrConditionalToNodeIDRequired   = errors.NewError("NODE-014", "Conditional ToNodeID is required")
	ErrInvalidNextID                 = errors.NewError("NODE-015", "This node kind cannot have NextID")
)
