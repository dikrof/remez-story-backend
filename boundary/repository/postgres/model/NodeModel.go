package postgresModel

import (
	"database/sql"
	"encoding/json"

	"remez_story/domain/entity/chapter"
	"remez_story/domain/entity/event"
	"remez_story/domain/entity/node"
	"remez_story/infrastructure/errors"
)

type NodeModel struct {
	ID              int64          `db:"id"`
	ChapterID       int64          `db:"chapter_id"`
	SceneLabel      string         `db:"scene_label"`
	Kind            string         `db:"kind"`
	Speaker         sql.NullString `db:"speaker"`
	Text            sql.NullString `db:"text"`
	NextID          sql.NullInt64  `db:"next_id"`
	ChoicesJSON     []byte         `db:"choices"`
	ConditionalJSON []byte         `db:"conditional"`
}

type ChoiceJSON struct {
	ID       int64      `json:"id"`
	Text     string     `json:"text"`
	Effects  EffectJSON `json:"effects"`
	ToNodeID int64      `json:"to_node_id"`
}

type EffectJSON struct {
	Add    []int64 `json:"add,omitempty"`
	Remove []int64 `json:"remove,omitempty"`
}

type ConditionalEdgeJSON struct {
	When struct {
		RequireAll  []int64 `json:"require_all,omitempty"`
		RequireNone []int64 `json:"require_none,omitempty"`
	} `json:"when"`
	ToNodeID int64 `json:"to_node_id"`
}

func NodeFromModel(model *NodeModel) (*node.Node, error) {
	errs := errors.NewErrors()

	nodeID, err := node.NewNodeID(model.ID)
	if err != nil {
		errs.AddError(err)
	}

	chapterID, err := chapter.NewChapterID(model.ChapterID)
	if err != nil {
		errs.AddError(err)
	}

	sceneLabel, err := node.NewSceneLabel(model.SceneLabel)
	if err != nil {
		errs.AddError(err)
	}

	kind, err := node.ParseNodeKind(model.Kind)
	if err != nil {
		errs.AddError(err)
	}

	if errs.IsPresent() {
		return nil, errs
	}

	builder := node.NewNodeBuilder().
		ID(nodeID).
		ChapterID(chapterID).
		SceneLabelValue(sceneLabel).
		Kind(kind)

	if model.Speaker.Valid {
		builder.Speaker(model.Speaker.String)
	}
	if model.Text.Valid {
		builder.Text(model.Text.String)
	}
	if model.NextID.Valid {
		nextID, err := node.NewNodeID(model.NextID.Int64)
		if err != nil {
			return nil, err
		}
		builder.NextID(nextID)
	}

	if len(model.ChoicesJSON) > 0 {
		var choicesJSON []ChoiceJSON
		if err := json.Unmarshal(model.ChoicesJSON, &choicesJSON); err != nil {
			return nil, err
		}

		choices := make([]node.Choice, 0, len(choicesJSON))
		for _, cj := range choicesJSON {
			choiceID, err := node.NewChoiceID(cj.ID)
			if err != nil {
				return nil, err
			}
			toNodeID, err := node.NewNodeID(cj.ToNodeID)
			if err != nil {
				return nil, err
			}

			addEvents := make([]event.EventID, 0, len(cj.Effects.Add))
			for _, eventIDInt := range cj.Effects.Add {
				eventID, err := event.NewEventID(eventIDInt)
				if err != nil {
					return nil, err
				}
				addEvents = append(addEvents, eventID)
			}

			removeEvents := make([]event.EventID, 0, len(cj.Effects.Remove))
			for _, eventIDInt := range cj.Effects.Remove {
				eventID, err := event.NewEventID(eventIDInt)
				if err != nil {
					return nil, err
				}
				removeEvents = append(removeEvents, eventID)
			}

			choice := node.Choice{
				ID:   choiceID,
				Text: cj.Text,
				Effects: []node.Effect{
					{
						Add:    addEvents,
						Remove: removeEvents,
					},
				},
				ToNodeID: toNodeID,
			}
			choices = append(choices, choice)
		}

		builder.Choices(choices)
	}

	if len(model.ConditionalJSON) > 0 {
		var conditionalJSON []ConditionalEdgeJSON
		if err := json.Unmarshal(model.ConditionalJSON, &conditionalJSON); err != nil {
			return nil, err
		}

		conditional := make([]node.ConditionalEdge, 0, len(conditionalJSON))
		for _, cej := range conditionalJSON {
			toNodeID, err := node.NewNodeID(cej.ToNodeID)
			if err != nil {
				return nil, err
			}

			requireAll := make([]event.EventID, 0, len(cej.When.RequireAll))
			for _, eventIDInt := range cej.When.RequireAll {
				eventID, err := event.NewEventID(eventIDInt)
				if err != nil {
					return nil, err
				}
				requireAll = append(requireAll, eventID)
			}

			requireNone := make([]event.EventID, 0, len(cej.When.RequireNone))
			for _, eventIDInt := range cej.When.RequireNone {
				eventID, err := event.NewEventID(eventIDInt)
				if err != nil {
					return nil, err
				}
				requireNone = append(requireNone, eventID)
			}

			edge := node.ConditionalEdge{
				When: node.Condition{
					RequireAll:  requireAll,
					RequireNone: requireNone,
				},
				ToNodeID: toNodeID,
			}
			conditional = append(conditional, edge)
		}

		builder.Conditional(conditional)
	}

	return builder.Build()
}
