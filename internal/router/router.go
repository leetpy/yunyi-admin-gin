package router

import (
	v1 "server/internal/router/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	apiv1 := r.Group("/api/v1")
	v1.RegisterCaptcha(apiv1)
}
