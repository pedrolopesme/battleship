package game

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pedrolopesme/citta-server/core/domain"
	"github.com/pedrolopesme/citta-server/core/ports"
)

type gameService struct {
	repository ports.GameRepository
}

func New(repo ports.GameRepository) *gameService {
	return &gameService{
		repository: repo,
	}
}

func (srv *gameService) Get(id string) (*domain.Game, error) {
	game, err := srv.repository.Get(id)
	if err != nil {
		return nil, errors.New("get game from repository has failed")
	}
	return game, nil
}

func (srv *gameService) Create(cols, rows uint) (*domain.Game, error) {
	gm := &domain.Game{
		ID: uuid.New().String(),
	}

	if err := srv.repository.Save(*gm); err != nil {
		return nil, errors.New("impossible to save game")
	}

	return gm, nil
}
