package dto

type StartGameRequest struct {
	PlayerID  string `json:"player_id"`
	ChapterID *int64 `json:"chapter_id"`
}

type MakeChoiceRequest struct {
	PlayerID string `json:"player_id"`
	ChoiceID int64  `json:"choice_id"`
}

type NavigationResponse struct {
	CurrentNode *NodeDetailDTO   `json:"current_node"`
	NextNodes   []*NodeDetailDTO `json:"next_nodes"`
	PlayerState *PlayerStateDTO  `json:"player_state"`
	ChapterInfo *ChapterInfoDTO  `json:"chapter_info"`
}

type NodeDetailDTO struct {
	ID          int64                 `json:"id"`
	ChapterID   int64                 `json:"chapter_id"`
	SceneLabel  string                `json:"scene_label"`
	Kind        string                `json:"kind"`
	Speaker     string                `json:"speaker,omitempty"`
	Text        string                `json:"text"`
	Choices     []*ChoiceDTO          `json:"choices,omitempty"`
	Conditional []*ConditionalEdgeDTO `json:"conditional,omitempty"`
	NextID      *int64                `json:"next_id,omitempty"`
}

type ChoiceDTO struct {
	ID       int64     `json:"id"`
	Text     string    `json:"text"`
	Effects  EffectDTO `json:"effects"`
	ToNodeID int64     `json:"to_node_id"`
}

type EffectDTO struct {
	Add        []int64               `json:"add,omitempty"`
	Remove     []int64               `json:"remove,omitempty"`
	MoneyDelta int                   `json:"money_delta,omitempty"`
	Relations  []ReputationChangeDTO `json:"relations,omitempty"`
}

type ReputationChangeDTO struct {
	Character string `json:"character"`
	Delta     int    `json:"delta"`
}

type ConditionalEdgeDTO struct {
	RequireAll  []int64 `json:"require_all,omitempty"`
	RequireNone []int64 `json:"require_none,omitempty"`
	ToNodeID    int64   `json:"to_node_id"`
}

type PlayerStateDTO struct {
	PlayerID       string         `json:"player_id"`
	ActiveEvents   []int64        `json:"active_events"`
	DecisionsCount int            `json:"decisions_count"`
	CurrentChapter int            `json:"current_chapter"`
	Money          int            `json:"money"`
	Reputation     map[string]int `json:"reputation"`
}

type ChapterInfoDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OrderIndex  int    `json:"order_index"`
}
