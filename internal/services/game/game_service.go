package game

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pedrolopesme/battleship/internal/domain"
	"github.com/pedrolopesme/battleship/internal/ports"
)

const (
	BOARD_ROWS = 10
	BOARD_COLS = 10
)

type gameService struct {
	repository ports.MatchRepository
}

func New(repo ports.MatchRepository) *gameService {
	return &gameService{
		repository: repo,
	}
}

func (srv *gameService) Get(id string) (*domain.Match, error) {
	game, err := srv.repository.Get(id)
	if err != nil {
		return nil, errors.New("get match from repository has failed")
	}
	return game, nil
}

func (srv *gameService) Create() (*domain.Match, error) {
	board := domain.Board{
		Settings: domain.BoardSettings{
			Cols: BOARD_COLS,
			Rows: BOARD_ROWS,
		},
	}

	match := domain.Match{
		ID:    uuid.New().String(),
		State: domain.GAME_WAITING,
		Board: board,
	}

	if err := srv.repository.Save(match); err != nil {
		return nil, errors.New("impossible to save match")
	}

	return &match, nil
}
