package node

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"remez_story/domain/entity/chapter"
)

type Node struct {
	ID         NodeID
	ChapterID  chapter.ChapterID
	SceneLabel SceneLabel
	Kind       NodeKind
	Speaker    string
	Text       string

	NextID *NodeID

	Choices     []Choice
	Conditional []ConditionalEdge

	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n *Node) SetNext(nextID NodeID) {
	n.NextID = &nextID
}

func (n *Node) ClearNext() {
	n.NextID = nil
}

func (n *Node) HasNext() bool {
	return n.NextID != nil
}

func (n *Node) GetNext() (NodeID, bool) {
	if n.NextID == nil {
		return NodeID{}, false
	}
	return *n.NextID, true
}

func (n *Node) Validate() error {
	if n.ID.IsZero() {
		return errors.New("node ID is required")
	}

	if !n.Kind.IsValid() {
		return errors.New("invalid node kind")
	}

	switch n.Kind {
	case NodeChoice:
		if len(n.Choices) == 0 {
			return errors.New("choice node must have choices")
		}
		if n.NextID != nil {
			return errors.New("choice node cannot have NextID - uses Choices instead")
		}
		for i, ch := range n.Choices {
			if ch.ToNodeID.IsZero() {
				return fmt.Errorf("choice[%d]: ToNodeID is required", i)
			}
			if strings.TrimSpace(ch.Text) == "" {
				return fmt.Errorf("choice[%d]: text is required", i)
			}
		}

	case NodeChoiceOption:
		if strings.TrimSpace(n.Text) == "" {
			return errors.New("choice option must have text")
		}

	case NodeNarration, NodeDialogue:
		if strings.TrimSpace(n.Text) == "" {
			return errors.New("text is required")
		}

	case NodeSystemNotification:
		if strings.TrimSpace(n.Text) == "" {
			return errors.New("notification text is required")
		}

		if n.NextID != nil {
			return errors.New("system notification should not have NextID")
		}
	}

	for i, edge := range n.Conditional {
		if edge.ToNodeID.IsZero() {
			return fmt.Errorf("conditional[%d]: ToNodeID is required", i)
		}
	}

	return nil
}
