package api

import (
	"net/http"

	gameApi "github.com/pedrolopesme/citta-server/api/game"
)

func SetupRoutes() {
	http.HandleFunc("/game/ws", gameApi.NewGameWebsocket().Run)
}
