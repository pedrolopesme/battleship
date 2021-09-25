package ports

import "github.com/pedrolopesme/battleship/internal/domain"

type MatchRepository interface {
	Get(id string) (*domain.Match, error)
	Save(domain.Match) error
}
