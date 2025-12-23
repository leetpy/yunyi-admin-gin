package response

import "github.com/gin-gonic/gin"

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Fail(c *gin.Context, err error) {}

func Success(c *gin.Context, data any) {}
