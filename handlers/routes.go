package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	game "github.com/pedrolopesme/citta-server/handlers/game"
	"github.com/pedrolopesme/citta-server/internal/ports"
)

func SetupRoutes(gamesService ports.GameService) {
	r := mux.NewRouter()
	r.HandleFunc("/game/ws", game.NewGameWebsocket().Run)
	r.HandleFunc("/game", game.NewPagesHandler(gamesService).HomePage)
	http.Handle("/", r)
}
