package cmd

import (
	"net/http"

	"github.com/pedrolopesme/citta-server/handlers"
	gamePorts "github.com/pedrolopesme/citta-server/internal/ports"
	gameRepository "github.com/pedrolopesme/citta-server/internal/repositories/game"
	gameService "github.com/pedrolopesme/citta-server/internal/services/game"
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
	logger      *zap.Logger
	gameService gamePorts.GameService
}

func NewBattleshipServer() *BattleshipServer {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	gameRepo := gameRepository.NewMemory()
	service := gameService.New(gameRepo)

	return &BattleshipServer{
		logger:      logger,
		gameService: service,
	}
}

func (b *BattleshipServer) Run() error {
	handlers.SetupRoutes(b.gameService)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		b.logger.Error(err.Error())
		return err
	}
	return nil
}
