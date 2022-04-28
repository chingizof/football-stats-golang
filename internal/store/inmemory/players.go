package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/chingizof/football-stats-golang/internal/models"
)

type PlayersRepo struct {
	data map[int]*models.Player

	mu *sync.RWMutex
}

func (db *PlayersRepo) Create(ctx context.Context, player *models.Player) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[player.ID] = player
	return nil
}

func (db *PlayersRepo) All(ctx context.Context) ([]*models.Player, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	players := make([]*models.Player, 0, len(db.data))
	for _, player := range db.data {
		players = append(players, player)
	}

	return players, nil
}

func (db *PlayersRepo) ByID(ctx context.Context, id int) (*models.Player, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	player, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no player with id %d", id)
	}

	return player, nil
}

func (db *PlayersRepo) Update(ctx context.Context, player *models.Player) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[player.ID] = player
	return nil
}

func (db *PlayersRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
