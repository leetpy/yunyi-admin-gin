package infra

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		uni := ut.New(zhT, zhT)

		Trans, _ = uni.GetTranslator("zh")
		_ = zhTranslations.RegisterDefaultTranslations(v, Trans)

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			if name == "" {
				return fld.Name
			}
			return name
		})
	}
}
