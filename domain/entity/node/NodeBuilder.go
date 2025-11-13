package node

import (
	"remez_story/domain/entity/chapter"
	"remez_story/infrastructure/errors"
	commonTime "remez_story/infrastructure/tools/time"
)

type NodeBuilder struct {
	node   *Node
	errors *errors.Errors
}

func NewNodeBuilder() *NodeBuilder {
	return &NodeBuilder{
		node:   &Node{},
		errors: errors.NewErrors(),
	}
}

func (b *NodeBuilder) ID(id NodeID) *NodeBuilder {
	b.node.ID = id
	return b
}

func (b *NodeBuilder) ChapterID(chapterID chapter.ChapterID) *NodeBuilder {
	b.node.ChapterID = chapterID
	return b
}

func (b *NodeBuilder) SceneLabel(label SceneLabel) *NodeBuilder {
	b.node.SceneLabel = label
	return b
}

func (b *NodeBuilder) Kind(kind NodeKind) *NodeBuilder {
	b.node.Kind = kind
	return b
}

func (b *NodeBuilder) Speaker(speaker string) *NodeBuilder {
	b.node.Speaker = speaker
	return b
}

func (b *NodeBuilder) Text(text string) *NodeBuilder {
	b.node.Text = text
	return b
}

func (b *NodeBuilder) NextID(nextID NodeID) *NodeBuilder {
	b.node.NextID = &nextID
	return b
}

func (b *NodeBuilder) Choices(choices []Choice) *NodeBuilder {
	b.node.Choices = choices
	return b
}

func (b *NodeBuilder) Conditional(conditional []ConditionalEdge) *NodeBuilder {
	b.node.Conditional = conditional
	return b
}

func (b *NodeBuilder) Version(version int) *NodeBuilder {
	b.node.Version = version
	return b
}

func (b *NodeBuilder) CreatedAt(createdAt *commonTime.Time) *NodeBuilder {
	b.node.CreatedAt = createdAt
	return b
}

func (b *NodeBuilder) UpdatedAt(updatedAt *commonTime.Time) *NodeBuilder {
	b.node.UpdatedAt = updatedAt
	return b
}

func (b *NodeBuilder) Build() (*Node, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()
	return b.node, nil
}

func (b *NodeBuilder) checkRequiredFields() {
	if b.node.ID.IsZero() {
		b.errors.AddError(ErrNodeIDRequired)
	}

	if b.node.ChapterID.IsZero() {
		b.errors.AddError(ErrChapterIDRequired)
	}

	if !b.node.Kind.IsValid() {
		b.errors.AddError(ErrInvalidNodeKind)
	}

	if b.node.Kind.RequiresText() && b.node.Text == "" {
		b.errors.AddError(ErrTextRequired)
	}

	if b.node.Kind.MustHaveChoices() && len(b.node.Choices) == 0 {
		b.errors.AddError(ErrChoicesRequired)
	}

	if !b.node.Kind.CanHaveNext() && b.node.NextID != nil {
		b.errors.AddError(ErrInvalidNextID)
	}
}

func (b *NodeBuilder) fillDefaultFields() {
	now := commonTime.Now()

	if b.node.CreatedAt == nil || b.node.CreatedAt.IsZero() {
		b.node.CreatedAt = now
	}

	if b.node.UpdatedAt == nil || b.node.UpdatedAt.IsZero() {
		b.node.UpdatedAt = now
	}

	if b.node.Version == 0 {
		b.node.Version = 1
	}
}
