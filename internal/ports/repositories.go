package ports

import "github.com/pedrolopesme/citta-server/internal/domain"

type GameRepository interface {
	Get(id string) (*domain.Game, error)
	Save(domain.Game) error
}
