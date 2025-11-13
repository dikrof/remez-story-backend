package player

import (
	"remez_story/domain/entity/node"
	commonTime "remez_story/infrastructure/tools/time"
)

type DecisionRecord struct {
	NodeID   node.NodeID
	ChoiceID node.ChoiceID
	At       *commonTime.Time
}
