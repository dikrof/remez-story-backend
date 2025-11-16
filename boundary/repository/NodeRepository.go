package repositoryInterface

import (
	"context"
	"remez_story/domain/entity/chapter"
	"remez_story/domain/entity/event"
	"remez_story/domain/entity/node"
)

type NodeRepository interface {
	GetByID(ctx context.Context, id node.NodeID) (*node.Node, error)
	GetByIDs(ctx context.Context, ids []node.NodeID) ([]*node.Node, error)
	GetStartNodeForChapter(ctx context.Context, chapterID chapter.ChapterID) (*node.Node, error)
	GetNextPossibleNodes(ctx context.Context, currentNodeID node.NodeID, playerState map[event.EventID]struct{}) ([]*node.Node, error)
}
