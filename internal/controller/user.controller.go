package controller

import (
	"github.com/carmenphan/go-ecommerce-backend-api/internal/service"
	"github.com/carmenphan/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 2001, []string{"tipjs", "m10"})
}
