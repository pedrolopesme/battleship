package game

import (
	"net/http"

	"github.com/pedrolopesme/citta-server/internal/ports"
)

type PagesHandler struct {
	gamesService ports.GameService
}

func NewPagesHandler(gamesService ports.GameService) *PagesHandler {
	return &PagesHandler{
		gamesService: gamesService,
	}
}

func (hdl *PagesHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/game.html")
}
