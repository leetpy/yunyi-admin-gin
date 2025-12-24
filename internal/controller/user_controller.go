package controller

import (
	"server/internal/dto"
	"server/internal/pkg/response"
	"server/internal/service"

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

func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}

	info, err := uc.userService.CreateUser(req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, gin.H{
		"id": info.ID,
	})
}
