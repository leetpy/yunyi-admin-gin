package v1

import (
	"server/internal/controller"

	"github.com/gin-gonic/gin"
)

// 添加用户
func CreateUser(r *gin.RouterGroup) {
	r.POST("/user", controller.CreateUser)
}
