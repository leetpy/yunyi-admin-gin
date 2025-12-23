package config

type AppConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
