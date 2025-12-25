package bootstrap

import "server/internal/infra"

func InitRds() error {
	return infra.InitRedis(infra.AppConfig.Redis)
}
