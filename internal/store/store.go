package store

import (
	"context"

	"github.com/chingizof/football-stats-golang/internal/models"
)

type Store interface {
	Players() PlayersRepository
	Coaches() CoachesRepository
}

type PlayersRepository interface {
	Create(ctx context.Context, player *models.Player) error
	All(ctx context.Context) ([]*models.Player, error)
	ByID(ctx context.Context, id int) (*models.Player, error)
	Update(ctx context.Context, player *models.Player) error
	Delete(ctx context.Context, id int) error
}

type CoachesRepository interface {
	Create(ctx context.Context, coach *models.Coach) error
	All(ctx context.Context) ([]*models.Coach, error)
	ByID(ctx context.Context, id int) (*models.Coach, error)
	Update(ctx context.Context, coach *models.Coach) error
	Delete(ctx context.Context, id int) error
}
