package player

import (
	"remez_story/domain/entity/node"
	commonTime "remez_story/infrastructure/tools/time"
)

type Progress struct {
	PlayerID      PlayerID
	CurrentNodeID *node.NodeID
	State         State
	Decisions     []DecisionRecord
	StartedAt     *commonTime.Time
	UpdatedAt     *commonTime.Time
}

func (p *Progress) Reset(to node.NodeID) {
	p.CurrentNodeID = &to
	p.State = NewState()
	p.Decisions = nil
	now := commonTime.Now()
	p.StartedAt = now
	p.UpdatedAt = now
}
