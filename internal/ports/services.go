package ports

import "github.com/pedrolopesme/battleship/internal/domain"

type GameService interface {
	Get(id string) (*domain.Game, error)
	Create(cols, rows uint) (*domain.Game, error)
}
