package infra

import (
	"context"
	"server/internal/config"

	"github.com/redis/go-redis/v9"
)

var RDS redis.UniversalClient

func InitRedis(rdsCfg config.RedisConfig) error {
	RDS = redis.NewClient(&redis.Options{
		Addr:     rdsCfg.Addr,
		Password: rdsCfg.Password,
		DB:       rdsCfg.DB,
	})

	_, err := RDS.Ping(context.Background()).Result()
	return err
}
