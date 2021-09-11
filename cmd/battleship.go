package cmd

import (
	"net/http"

	api "github.com/pedrolopesme/citta-server/api"
	"go.uber.org/zap"
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
	logger *zap.Logger
}

func NewBattleshipServer() *BattleshipServer {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return &BattleshipServer{
		logger: logger,
	}
}

func (b *BattleshipServer) Run() error {
	api.SetupRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		b.logger.Error(err.Error())
		return err
	}
	return nil
}
