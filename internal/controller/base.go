package controller

import (
	"server/internal/infra"
	"server/internal/pkg/apperror"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseController struct{}

func (bc *BaseController) shouldBindJSON(c *gin.Context, req any) error {
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		msg := errs[0].Translate(infra.Trans) // 取第一个错误
		return apperror.New(apperror.InvalidParam, msg)
	}

	return nil
}
