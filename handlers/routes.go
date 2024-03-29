package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	game "github.com/pedrolopesme/battleship/handlers/game"
	gameWebsocket "github.com/pedrolopesme/battleship/handlers/game/websocket"
	"github.com/pedrolopesme/battleship/internal/ports"
)

func SetupRoutes(gamesService ports.GameService) {
	r := mux.NewRouter()
	r.HandleFunc("/game/ws", gameWebsocket.NewGameWebsocket(gamesService).Run)
	r.HandleFunc("/game", game.NewPagesHandler(gamesService).HomePage)

	fs := http.FileServer(http.Dir("./web/"))
	r.PathPrefix("/statics/").Handler(http.StripPrefix("/statics/", fs))

	http.Handle("/", r)
}
