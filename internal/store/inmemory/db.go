package inmemory

import (
	"sync"

	"github.com/chingizof/football-stats-golang/internal/models"
	"github.com/chingizof/football-stats-golang/internal/store"
)

type DB struct {
	playersRepo store.PlayersRepository
	coachesRepo store.CoachesRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Players() store.PlayersRepository {
	if db.playersRepo == nil {
		db.playersRepo = &PlayersRepo{
			data: make(map[int]*models.Player),
			mu:   new(sync.RWMutex),
		}
	}

	return db.playersRepo
}

func (db *DB) Coaches() store.CoachesRepository {
	if db.coachesRepo == nil {
		db.coachesRepo = &CoachesRepo{
			data: make(map[int]*models.Coach),
			mu:   new(sync.RWMutex),
		}
	}

	return db.coachesRepo
}
