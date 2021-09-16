package ports

import "github.com/pedrolopesme/battleship/internal/domain"

type GameRepository interface {
	Get(id string) (*domain.Game, error)
	Save(domain.Game) error
}
