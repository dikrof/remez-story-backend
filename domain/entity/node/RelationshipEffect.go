package node

import "remez_story/domain/entity/character"

type ReputationChange struct {
	Character character.CharacterCode `json:"character"`
	Delta     int                     `json:"delta"`
}
