package v1

import (
	"server/internal/controller"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 注册验证码
func RegisterUser(r *gin.RouterGroup) {
	userController := controller.NewUserController()

	// 登录
	r.POST("/login", userController.Login)

	// 添加用户
	userGroup := r.Group("/user", middleware.JwtAuth())
	userGroup.POST("", userController.CreateUser)
}
