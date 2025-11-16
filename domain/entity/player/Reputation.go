package player

import "remez_story/domain/entity/character"

type Reputation struct {
	Scores map[character.CharacterCode]int
}

func NewReputation() Reputation {
	return Reputation{Scores: map[character.CharacterCode]int{}}
}

func (r *Reputation) Get(code character.CharacterCode) int {
	if r.Scores == nil {
		r.Scores = map[character.CharacterCode]int{}
	}
	return r.Scores[code]
}

func (r *Reputation) Add(code character.CharacterCode, delta int) {
	if r.Scores == nil {
		r.Scores = map[character.CharacterCode]int{}
	}
	r.Scores[code] += delta
}
