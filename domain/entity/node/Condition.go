package node

import "remez_story/domain/entity/event"

type Condition struct {
	RequireAll  []event.EventID `json:"require_all,omitempty"`
	RequireNone []event.EventID `json:"require_none,omitempty"`
}

func (c Condition) IsSatisfied(state map[event.EventID]struct{}) bool {
	for _, eventID := range c.RequireAll {
		if _, exists := state[eventID]; !exists {
			return false
		}
	}
	for _, eventID := range c.RequireNone {
		if _, exists := state[eventID]; exists {
			return false
		}
	}
	return true
}

type Effect struct {
	Add    []event.EventID `json:"add,omitempty"`
	Remove []event.EventID `json:"remove,omitempty"`

	MoneyDelta int                `json:"money_delta,omitempty"`
	Relations  []ReputationChange `json:"relations,omitempty"`
}

type ConditionalEdge struct {
	When     Condition `json:"when"`
	ToNodeID NodeID    `json:"to_node_id"`
}
