package ports

import "github.com/pedrolopesme/citta-server/core/domain"

type GameService interface {
	Get(id string) (*domain.Game, error)
	Create(cols, rows uint) (*domain.Game, error)
}
