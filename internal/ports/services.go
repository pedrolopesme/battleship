package ports

import "github.com/pedrolopesme/battleship/internal/domain"

type MatchService interface {
	Get(id string) (*domain.Match, error)
	Create() (*domain.Match, error)
}
