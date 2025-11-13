package node

import (
	"remez_story/domain/entity/chapter"
	"remez_story/infrastructure/errors"
	commonTime "remez_story/infrastructure/tools/time"
	"strings"
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
	CreatedAt *commonTime.Time
	UpdatedAt *commonTime.Time
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
	errs := errors.NewErrors()

	if n.ID.IsZero() {
		errs.AddError(ErrNodeIDRequired)
	}

	if !n.Kind.IsValid() {
		errs.AddError(ErrInvalidNodeKind)
	}

	if n.Kind.RequiresText() && strings.TrimSpace(n.Text) == "" {
		errs.AddError(ErrTextRequired)
	}

	if n.Kind.MustHaveChoices() {
		if len(n.Choices) == 0 {
			errs.AddError(ErrChoicesRequired)
		}

		for _, ch := range n.Choices {
			if ch.ToNodeID.IsZero() {
				errs.AddError(ErrChoiceToNodeIDRequired)
			}
			if strings.TrimSpace(ch.Text) == "" {
				errs.AddError(ErrChoiceTextIsRequired)
			}
		}
	}

	if !n.Kind.CanHaveNext() && n.NextID != nil {
		errs.AddError(ErrInvalidNextID)
	}

	for _, edge := range n.Conditional {
		if edge.ToNodeID.IsZero() {
			errs.AddError(ErrConditionalToNodeIDRequired)
		}
	}

	if errs.IsPresent() {
		return errs
	}

	return nil
}
