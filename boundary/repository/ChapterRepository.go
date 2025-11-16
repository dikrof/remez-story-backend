package repositoryInterface

import (
	"context"
	"remez_story/domain/entity/chapter"
)

type ChapterRepository interface {
	GetByID(ctx context.Context, id chapter.ChapterID) (*chapter.Chapter, error)
	GetByOrderIndex(ctx context.Context, orderIndex int) (*chapter.Chapter, error)
	FindAll(ctx context.Context) ([]*chapter.Chapter, error)
}
