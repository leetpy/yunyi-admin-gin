package bootstrap

import (
	"server/internal/infra"
	"server/internal/router"

	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	infra.InitValidator()
	r := gin.Default()
	router.RegisterRouter(&r.RouterGroup)
	return r
}
