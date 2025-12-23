package config

type Config struct {
	App     AppConfig     `mapstructure:"app"`
	Captcha CaptchaConfig `mapstructure:"captcha"`
}
