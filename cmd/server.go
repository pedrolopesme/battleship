package cmd

import (
	"github.com/gin-gonic/gin"
	gameService "github.com/pedrolopesme/citta-server/core/services/game"
	gameHandler "github.com/pedrolopesme/citta-server/handlers/game"
	gameRepository "github.com/pedrolopesme/citta-server/repositories/game"
)

func Server() {
	repo := gameRepository.NewMemory()
	service := gameService.New(repo)
	handler := gameHandler.NewHTTPHandler(service)

	router := gin.New()
	router.GET("/games/:id", handler.Get)

	router.Run(":8080")
}
