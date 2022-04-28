package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/chingizof/football-stats-golang/internal/models"
)

type CoachesRepo struct {
	data map[int]*models.Coach

	mu *sync.RWMutex
}

func (db *CoachesRepo) Create(ctx context.Context, phone *models.Coach) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[phone.ID] = phone
	return nil
}

func (db *CoachesRepo) All(ctx context.Context) ([]*models.Coach, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	coaches := make([]*models.Coach, 0, len(db.data))
	for _, coach := range db.data {
		coaches = append(coaches, coach)
	}

	return coaches, nil
}

func (db *CoachesRepo) ByID(ctx context.Context, id int) (*models.Coach, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	coach, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no coach with id %d", id)
	}

	return coach, nil
}

func (db *CoachesRepo) Update(ctx context.Context, coach *models.Coach) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[coach.ID] = coach
	return nil
}

func (db *CoachesRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
