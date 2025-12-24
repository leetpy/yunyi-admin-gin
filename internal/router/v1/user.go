package v1

import (
	"server/internal/controller"

	"github.com/gin-gonic/gin"
)

// 注册验证码
func RegisterUser(r *gin.RouterGroup) {
	userController := controller.NewUserController()

	// 添加用户
	r.POST("/user", userController.CreateUser)
}
