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

func (p *Progress) GetPlayerID() PlayerID {
	return p.PlayerID
}

func (p *Progress) GetCurrentNodeID() (node.NodeID, bool) {
	if p.CurrentNodeID == nil {
		return node.NodeID{}, false
	}
	return *p.CurrentNodeID, true
}

func (p *Progress) GetState() State {
	return p.State
}

func (p *Progress) MoveTo(nodeID node.NodeID) {
	p.CurrentNodeID = &nodeID
	p.UpdatedAt = commonTime.Now()
}

func (p *Progress) RecordDecision(nodeID node.NodeID, choiceID node.ChoiceID) {
	decision := DecisionRecord{
		NodeID:   nodeID,
		ChoiceID: choiceID,
		At:       commonTime.Now(),
	}
	p.Decisions = append(p.Decisions, decision)
	p.UpdatedAt = commonTime.Now()
}

func (p *Progress) ApplyEffects(effects node.Effect) {
	p.State.ApplyEffect(effects)
	p.UpdatedAt = commonTime.Now()
}

func (p *Progress) Reset(to node.NodeID) {
	p.CurrentNodeID = &to
	p.State = NewState()
	p.Decisions = nil
	now := commonTime.Now()
	p.StartedAt = now
	p.UpdatedAt = now
}
