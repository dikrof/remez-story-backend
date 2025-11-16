package repositoryInterface

import (
	"context"
	"remez_story/domain/entity/player"
)

type ProgressRepository interface {
	Save(ctx context.Context, progress *player.Progress) error
	GetByPlayerID(ctx context.Context, playerID player.PlayerID) (*player.Progress, error)
	Update(ctx context.Context, progress *player.Progress) error
	Exists(ctx context.Context, playerID player.PlayerID) (bool, error)
}
