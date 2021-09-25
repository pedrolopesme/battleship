package game

import (
	"net/http"

	"github.com/pedrolopesme/battleship/internal/ports"
)

type PagesHandler struct {
	gamesService ports.MatchService
}

func NewPagesHandler(gamesService ports.MatchService) *PagesHandler {
	return &PagesHandler{
		gamesService: gamesService,
	}
}

func (hdl *PagesHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/game.html")
}
