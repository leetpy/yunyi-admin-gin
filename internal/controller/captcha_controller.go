package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"server/internal/pkg/response"
	"server/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	captchaService *service.CaptchaService
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{
		captchaService: service.NewCaptchaService(),
	}
}

func (cc *CaptchaController) GetCaptcha(c *gin.Context) {
	captId, content, err := cc.captchaService.GetCaptcha()
	if err != nil {
		response.Fail(c, err)
		return
	}

	name := fmt.Sprintf("%s.png", captId)
	c.SetCookie("mission_mnt_captcha", captId, 600, "/", "", false, true)
	http.ServeContent(c.Writer, c.Request, name, time.Now(), bytes.NewReader(content))
}
