package player

import (
	"time"

	"remez_story/domain/entity/node"
	"remez_story/infrastructure/errors"
)

type ProgressBuilder struct {
	progress *Progress
	errors   *errors.Errors
}

func NewProgressBuilder() *ProgressBuilder {
	return &ProgressBuilder{
		progress: &Progress{},
		errors:   errors.NewErrors(),
	}
}

func (b *ProgressBuilder) PlayerID(playerID PlayerID) *ProgressBuilder {
	b.progress.PlayerID = playerID
	return b
}

func (b *ProgressBuilder) CurrentNodeID(currentNodeID node.NodeID) *ProgressBuilder {
	b.progress.CurrentNodeID = &currentNodeID
	return b
}

func (b *ProgressBuilder) State(state State) *ProgressBuilder {
	b.progress.State = state
	return b
}

func (b *ProgressBuilder) Decisions(decisions []DecisionRecord) *ProgressBuilder {
	b.progress.Decisions = decisions
	return b
}

func (b *ProgressBuilder) StartedAt(startedAt time.Time) *ProgressBuilder {
	b.progress.StartedAt = startedAt
	return b
}

func (b *ProgressBuilder) UpdatedAt(updatedAt time.Time) *ProgressBuilder {
	b.progress.UpdatedAt = updatedAt
	return b
}

func (b *ProgressBuilder) Build() (*Progress, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()
	return b.progress, nil
}

func (b *ProgressBuilder) checkRequiredFields() {
	if b.progress.PlayerID.IsZero() {
		b.errors.AddError(ErrPlayerIDRequired)
	}
}

func (b *ProgressBuilder) fillDefaultFields() {
	now := time.Now().UTC()

	if b.progress.StartedAt.IsZero() {
		b.progress.StartedAt = now
	}

	if b.progress.UpdatedAt.IsZero() {
		b.progress.UpdatedAt = now
	}

	if b.progress.State.Events == nil {
		b.progress.State = NewState()
	}

	if b.progress.Decisions == nil {
		b.progress.Decisions = []DecisionRecord{}
	}
}
