package game

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrolopesme/battleship/internal/ports"
)

type HTTPHandler struct {
	gamesService ports.GameService
}

func NewHTTPHandler(gamesService ports.GameService) *HTTPHandler {
	return &HTTPHandler{
		gamesService: gamesService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	game, err := hdl.gamesService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, game)
}
