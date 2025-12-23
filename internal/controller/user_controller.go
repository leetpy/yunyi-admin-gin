package controller

import (
	"server/internal/dto"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	var req dto.UserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, err)
		return
	}

	list := service.UserService.GetList(req)
	response.Success(c, list)
}
