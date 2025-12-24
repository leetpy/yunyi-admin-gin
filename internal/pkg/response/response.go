package response

import (
	"net/http"

	"server/internal/pkg/apperror"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Fail(c *gin.Context, err error) {
	if e, ok := err.(*apperror.BizError); ok {
		c.JSON(http.StatusOK, Result{
			Code: e.Code,
			Msg:  e.Message,
		})
		return
	}

	// 未知错误
	c.JSON(http.StatusOK, Result{
		Code: apperror.InternalError,
		Msg:  "系统错误",
	})
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}
