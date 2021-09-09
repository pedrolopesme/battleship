package cmd

import (
	"github.com/gin-gonic/gin"
	gameHandler "github.com/pedrolopesme/citta-server/cor"
	gameRepository "github.com/pedrolopesme/citta-server/core/repositories/game"
	gameService "github.com/pedrolopesme/citta-server/core/service/game"
)

func Server() {
	repo := gameRepository.NewMemory()
	gamesService := gameService.New(repo)
	gamesHandler := gameHandler.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}
