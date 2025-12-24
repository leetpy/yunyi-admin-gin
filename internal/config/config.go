package config

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Captcha  CaptchaConfig  `mapstructure:"captcha"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Zap      ZapConfig      `mapstructure:"zap"`
}
