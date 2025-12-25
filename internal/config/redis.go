package config

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`     // 服务器地址:端口
	Password string `mapstructure:"password"` // 密码
	DB       int    `mapstructure:"db"`       // 单实例模式下redis的哪个数据库
}
