package player

import (
	"time"

	"remez_story/domain/entity/node"
)

type DecisionRecord struct {
	NodeID   node.NodeID
	ChoiceID node.ChoiceID
	At       time.Time
}
