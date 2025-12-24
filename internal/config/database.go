package config

import (
	"fmt"
)

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Config       string `mapstructure:"config"`
	DBName       string `mapstructure:"db-name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxIdelConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap"`
}

func (c *DatabaseConfig) Dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.Config,
	)
}
