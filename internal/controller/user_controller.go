package controller

import (
	"server/internal/dto"
	"server/internal/pkg/apperror"
	"server/internal/pkg/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
	BaseController
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var req dto.UserLoginReq
	if err := uc.shouldBindJSON(c, &req); err != nil {
		response.Fail(c, apperror.New(apperror.InvalidParam, err.Error()))
		return
	}

	data, err := uc.userService.Login(req)
	if err != nil {
		response.Fail(c, apperror.New(apperror.LoginFailed, err.Error()))
		return
	}

	response.Success(c, data)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserReq
	if err := uc.shouldBindJSON(c, &req); err != nil {
		response.Fail(c, apperror.New(apperror.InvalidParam, err.Error()))
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
