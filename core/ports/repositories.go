package ports

import "github.com/pedrolopesme/citta-server/core/domain"

type GameRepository interface {
	Get(id string) (*domain.Game, error)
	Save(domain.Game) error
}
