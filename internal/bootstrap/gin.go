package bootstrap

import (
	"server/internal/router"

	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	r := gin.Default()
	router.RegisterRouter(&r.RouterGroup)
	return r
}
