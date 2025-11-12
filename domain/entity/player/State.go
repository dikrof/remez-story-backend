package player

import (
	"remez_story/domain/entity/event"
	"remez_story/domain/entity/node"
)

type State struct {
	Events map[event.EventID]struct{}
}

func NewState() State {
	return State{Events: map[event.EventID]struct{}{}}
}

func (s *State) Has(id event.EventID) bool {
	_, ok := s.Events[id]
	return ok
}

func (s *State) Add(id event.EventID) {
	s.Events[id] = struct{}{}
}

func (s *State) Remove(id event.EventID) {
	delete(s.Events, id)
}

func (s *State) ApplyEffect(e node.Effect) {
	for _, id := range e.Add {
		s.Add(id)
	}
	for _, id := range e.Remove {
		s.Remove(id)
	}
}
