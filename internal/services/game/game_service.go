package game

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pedrolopesme/citta-server/internal/domain"
	"github.com/pedrolopesme/citta-server/internal/ports"
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
	board := domain.Board{
		Settings: domain.BoardSettings{
			Cols: cols,
			Rows: rows,
		},
	}

	game := domain.Game{
		ID:    uuid.New().String(),
		State: domain.GAME_WAITING,
		Board: board,
	}

	if err := srv.repository.Save(game); err != nil {
		return nil, errors.New("impossible to save game")
	}

	return &game, nil
}
