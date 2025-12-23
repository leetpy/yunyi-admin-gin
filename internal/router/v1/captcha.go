package v1

import (
	"server/internal/controller"

	"github.com/gin-gonic/gin"
)

// 注册验证码
func RegisterCaptcha(r *gin.RouterGroup) {
	captchaController := controller.NewCaptchaController()

	// 生成验证码
	r.GET("/captcha", captchaController.GetCaptcha)
}
