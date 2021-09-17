package cmd

import (
	"net/http"

	"github.com/pedrolopesme/battleship/handlers"
	gamePorts "github.com/pedrolopesme/battleship/internal/ports"
	gameRepository "github.com/pedrolopesme/battleship/internal/repositories/game"
	gameService "github.com/pedrolopesme/battleship/internal/services/game"
	"go.uber.org/zap"
)

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
