package service

import (
	"image/color"
	"strings"

	"github.com/mojocn/base64Captcha"
)

var stroe = base64Captcha.DefaultMemStore

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

// 生成验证码
func (s *CaptchaService) GetCaptcha() (id string, result []byte, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          4,
		Source:          base64Captcha.TxtSimpleCharaters,
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, stroe)
	return s.Generate(captcha)
}

func (s *CaptchaService) Generate(c *base64Captcha.Captcha) (id string, result []byte, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	cItem := item.(*base64Captcha.ItemChar)
	if err != nil {
		return "", nil, err
	}
	err = c.Store.Set(id, answer)
	if err != nil {
		return "", nil, err
	}
	result = cItem.BinaryEncoding()
	return
}

// 验证验证码
func (s *CaptchaService) VerifyCaptcha(captcheId, capt string) bool {
	if len(strings.TrimSpace(captcheId)) > 0 && len(strings.TrimSpace(capt)) > 0 {
		return stroe.Verify(captcheId, capt, true)
	}
	return false
}
