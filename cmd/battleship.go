package cmd

import (
	"net/http"

	api "github.com/pedrolopesme/citta-server/api/game"
)

// func Server() {
// 	repo := gameRepository.NewMemory()
// 	service := gameService.New(repo)
// 	handler := gameHandler.NewHTTPHandler(service)

// 	router := gin.New()
// 	router.GET("/games/:id", handler.Get)

// 	router.Run(":8080")
// }

type BattleshipServer struct {
}

func NewBattleshipServer() *BattleshipServer {
	return &BattleshipServer{}
}

func (b *BattleshipServer) setupRoutes() {
	http.HandleFunc("/game/ws", api.NewGameWebsocket().Run)
}

func (b *BattleshipServer) Run() error {
	b.setupRoutes()

	http.ListenAndServe(":8080", nil)
	return nil
}
