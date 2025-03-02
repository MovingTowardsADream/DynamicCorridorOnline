package v1

import (
	"github.com/gin-gonic/gin"
)

type GameRoutes struct {
}

func newGameRoutes(handler *gin.RouterGroup) {
	r := &GameRoutes{}

	h := handler.Group("/game")
	{
		h.POST("", r.CreateGame)
	}
}

func (r *GameRoutes) CreateGame(c *gin.Context) {

}
