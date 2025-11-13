package node

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type NodeKind string

const (
	NodeNarration          NodeKind = "NARRATION"
	NodeDialogue           NodeKind = "DIALOGUE"
	NodeChoice             NodeKind = "CHOICE"
	NodeSystemNotification NodeKind = "SYSTEM_NOTIFICATION"
	NodeChoiceOption       NodeKind = "CHOICE_OPTION"
)

func ParseNodeKind(s string) (NodeKind, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	switch NodeKind(s) {
	case NodeNarration, NodeDialogue, NodeChoice, NodeSystemNotification, NodeChoiceOption:
		return NodeKind(s), nil
	default:
		return "", ErrInvalidNodeKind
	}
}

func (k NodeKind) IsValid() bool {
	_, err := ParseNodeKind(string(k))
	return err == nil
}

func (k NodeKind) RequiresText() bool {
	return k == NodeNarration ||
		k == NodeDialogue ||
		k == NodeSystemNotification ||
		k == NodeChoiceOption
}

func (k NodeKind) CanHaveNext() bool {
	return k != NodeChoice && k != NodeSystemNotification
}

func (k NodeKind) MustHaveChoices() bool {
	return k == NodeChoice
}

func (k NodeKind) CanHaveSpeaker() bool {
	return k == NodeDialogue
}

func (k NodeKind) IsInteractive() bool {
	return k == NodeChoice
}

func (k NodeKind) IsSystemGenerated() bool {
	return k == NodeSystemNotification
}

func (k NodeKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(k))
}

func (k *NodeKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	val, err := ParseNodeKind(s)
	if err != nil {
		return err
	}
	*k = val
	return nil
}

func (k NodeKind) String() string {
	return string(k)
}

func (k NodeKind) Value() (driver.Value, error) {
	return string(k), nil
}

func (k *NodeKind) Scan(src any) error {
	if src == nil {
		*k = ""
		return nil
	}

	switch t := src.(type) {
	case string:
		val, err := ParseNodeKind(t)
		if err != nil {
			return err
		}
		*k = val
		return nil
	case []byte:
		val, err := ParseNodeKind(string(t))
		if err != nil {
			return err
		}
		*k = val
		return nil
	default:
		return ErrNodeKindUnsupportedScanType
	}
}
