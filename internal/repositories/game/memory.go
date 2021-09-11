package game

import (
	"encoding/json"
	"errors"

	"github.com/pedrolopesme/citta-server/internal/domain"
)

type memory struct {
	kvs map[string][]byte
}

func NewMemory() *memory {
	return &memory{kvs: map[string][]byte{}}
}

func (repo *memory) Get(id string) (*domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return nil, errors.New("fail to get value from kvs")
		}

		return &game, nil
	}

	return nil, errors.New("game not found in kvs")
}

func (repo *memory) Save(game domain.Game) error {
	gameMarshaled, err := json.Marshal(game)
	if err != nil {
		return errors.New("impossible to store game")
	}

	repo.kvs[game.ID] = gameMarshaled
	return nil
}
