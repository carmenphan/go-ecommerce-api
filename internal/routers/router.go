package routers

import (
	"github.com/carmenphan/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}

	return r

}
func Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
	})
}
