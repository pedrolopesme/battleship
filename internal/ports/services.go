package ports

import "github.com/pedrolopesme/battleship/internal/domain"

type GameService interface {
	Get(id string) (*domain.Game, error)
	Create() (*domain.Game, error)
}
