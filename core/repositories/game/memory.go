package game

import (
	"encoding/json"
	"errors"

	"github.com/pedrolopesme/citta-server/core/domain"
)

type memory struct {
	kvs map[string][]byte
}

func NewMemory() *memory {
	return &memory{kvs: map[string][]byte{}}
}

func (repo *memory) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New("fail to get value from kvs")
		}

		return game, nil
	}

	return domain.Game{}, errors.New("game not found in kvs")
}
