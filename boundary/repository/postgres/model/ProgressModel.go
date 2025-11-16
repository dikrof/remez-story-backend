package postgresModel

import (
	"database/sql"
	"encoding/json"
	"remez_story/domain/entity/character"
	"remez_story/domain/entity/event"
	"remez_story/domain/entity/node"
	"remez_story/domain/entity/player"
	"remez_story/infrastructure/errors"
	commonTime "remez_story/infrastructure/tools/time"
)

type ProgressModel struct {
	PlayerID       string        `db:"player_id"`
	CurrentNodeID  sql.NullInt64 `db:"current_node_id"`
	StateJSON      []byte        `db:"state"`
	DecisionsJSON  []byte        `db:"decisions"`
	Money          int64         `db:"money"`
	ReputationJSON []byte        `db:"reputation"`
	StartedAt      int64         `db:"started_at"`
	UpdatedAt      int64         `db:"updated_at"`
}

type StateJSON struct {
	Events []int64 `json:"events"`
}

type DecisionJSON struct {
	NodeID   int64 `json:"node_id"`
	ChoiceID int64 `json:"choice_id"`
	At       int64 `json:"at"`
}

type ReputationJSON struct {
	Scores map[string]int `json:"scores"`
}

func ProgressToModel(p *player.Progress) (*ProgressModel, error) {
	model := &ProgressModel{
		PlayerID:  p.PlayerID.String(),
		Money:     int64(p.Money),
		StartedAt: p.StartedAt.UnixNano(),
		UpdatedAt: p.UpdatedAt.UnixNano(),
	}

	if currentNodeID, ok := p.GetCurrentNodeID(); ok {
		model.CurrentNodeID = sql.NullInt64{Int64: currentNodeID.Int64(), Valid: true}
	}

	state := p.GetState()
	eventIDs := make([]int64, 0, len(state.Events))
	for eventID := range state.Events {
		eventIDs = append(eventIDs, eventID.Int64())
	}
	stateJSON := StateJSON{Events: eventIDs}
	stateBytes, err := json.Marshal(stateJSON)
	if err != nil {
		return nil, err
	}
	model.StateJSON = stateBytes

	decisionsJSON := make([]DecisionJSON, 0, len(p.Decisions))
	for _, decision := range p.Decisions {
		decisionsJSON = append(decisionsJSON, DecisionJSON{
			NodeID:   decision.NodeID.Int64(),
			ChoiceID: decision.ChoiceID.Int64(),
			At:       decision.At.UnixNano(),
		})
	}
	decisionsBytes, err := json.Marshal(decisionsJSON)
	if err != nil {
		return nil, err
	}
	model.DecisionsJSON = decisionsBytes

	reputationScores := make(map[string]int)
	if p.Reputation.Scores != nil {
		for charCode, score := range p.Reputation.Scores {
			reputationScores[charCode.String()] = score
		}
	}
	reputationJSON := ReputationJSON{Scores: reputationScores}
	reputationBytes, err := json.Marshal(reputationJSON)
	if err != nil {
		return nil, err
	}
	model.ReputationJSON = reputationBytes

	return model, nil
}

func ProgressFromModel(model *ProgressModel) (*player.Progress, error) {
	errs := errors.NewErrors()

	playerID, err := player.NewPlayerID(model.PlayerID)
	if err != nil {
		errs.AddError(err)
	}

	if errs.IsPresent() {
		return nil, errs
	}

	builder := player.NewProgressBuilder().
		PlayerID(playerID).
		Money(int(model.Money)).
		StartedAt(commonTime.FromUnixNano(model.StartedAt)).
		UpdatedAt(commonTime.FromUnixNano(model.UpdatedAt))

	if model.CurrentNodeID.Valid {
		currentNodeID, err := node.NewNodeID(model.CurrentNodeID.Int64)
		if err != nil {
			return nil, err
		}
		builder.CurrentNodeID(currentNodeID)
	}

	var stateJSON StateJSON
	if err := json.Unmarshal(model.StateJSON, &stateJSON); err != nil {
		return nil, err
	}

	state := player.NewState()
	for _, eventIDInt := range stateJSON.Events {
		eventID, err := event.NewEventID(eventIDInt)
		if err != nil {
			return nil, err
		}
		state.Add(eventID)
	}
	builder.State(state)

	var decisionsJSON []DecisionJSON
	if len(model.DecisionsJSON) > 0 {
		if err := json.Unmarshal(model.DecisionsJSON, &decisionsJSON); err != nil {
			return nil, err
		}

		decisions := make([]player.DecisionRecord, 0, len(decisionsJSON))
		for _, dj := range decisionsJSON {
			nodeID, err := node.NewNodeID(dj.NodeID)
			if err != nil {
				return nil, err
			}

			choiceID, err := node.NewChoiceID(dj.ChoiceID)
			if err != nil {
				return nil, err
			}

			decision := player.DecisionRecord{
				NodeID:   nodeID,
				ChoiceID: choiceID,
				At:       commonTime.FromUnixNano(dj.At),
			}
			decisions = append(decisions, decision)
		}

		builder.Decisions(decisions)
	}

	var reputationJSON ReputationJSON
	if err := json.Unmarshal(model.ReputationJSON, &reputationJSON); err != nil {
		return nil, err
	}

	reputation := player.NewReputation()
	for charCodeStr, score := range reputationJSON.Scores {
		charCode, err := character.NewCharacterCode(charCodeStr)
		if err != nil {
			return nil, err
		}
		reputation.Scores[charCode] = score
	}
	builder.Reputation(reputation)

	return builder.Build()
}
