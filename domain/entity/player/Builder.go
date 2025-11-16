package player

import (
	"remez_story/domain/entity/node"
	"remez_story/infrastructure/errors"
	commonTime "remez_story/infrastructure/tools/time"
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

func (b *ProgressBuilder) Money(money int) *ProgressBuilder {
	b.progress.Money = money
	return b
}

func (b *ProgressBuilder) Reputation(reputation Reputation) *ProgressBuilder {
	b.progress.Reputation = reputation
	return b
}

func (b *ProgressBuilder) StartedAt(startedAt *commonTime.Time) *ProgressBuilder {
	b.progress.StartedAt = startedAt
	return b
}

func (b *ProgressBuilder) UpdatedAt(updatedAt *commonTime.Time) *ProgressBuilder {
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
	now := commonTime.Now()

	if b.progress.StartedAt == nil || b.progress.StartedAt.IsZero() {
		b.progress.StartedAt = now
	}

	if b.progress.UpdatedAt == nil || b.progress.UpdatedAt.IsZero() {
		b.progress.UpdatedAt = now
	}

	if b.progress.State.Events == nil {
		b.progress.State = NewState()
	}

	if b.progress.Decisions == nil {
		b.progress.Decisions = []DecisionRecord{}
	}

	if b.progress.Reputation.Scores == nil {
		b.progress.Reputation = NewReputation()
	}
}
