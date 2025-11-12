package player

import (
	"time"

	"remez_story/domain/entity/node"
)

type Progress struct {
	PlayerID      PlayerID
	CurrentNodeID *node.NodeID
	State         State
	Decisions     []DecisionRecord
	StartedAt     time.Time
	UpdatedAt     time.Time
}

func (p *Progress) Reset(to node.NodeID) {
	p.CurrentNodeID = &to
	p.State = NewState()
	p.Decisions = nil
	now := time.Now().UTC()
	p.StartedAt = now
	p.UpdatedAt = now
}
