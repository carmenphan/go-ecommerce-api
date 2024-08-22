package controller

import (
	"github.com/carmenphan/go-ecommerce-backend-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct {
	pongService *service.PongService
}

func NewPongController() *PongController {
	return &PongController{
		pongService: service.NewPongService(),
	}
}

func (pc *PongController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": pc.pongService.GetPong(),
	})
}
