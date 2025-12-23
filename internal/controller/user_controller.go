package controller

import (
	"server/internal/dto"
	"server/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// func GetUserList(c *gin.Context) {
// 	var req dto.UserListReq
// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		response.Fail(c, err)
// 		return
// 	}

// 	list := service.UserService.FindList(req)
// 	response.Success(c, list)
// }

func CreateUser(c *gin.Context) {
	var req dto.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}

}
