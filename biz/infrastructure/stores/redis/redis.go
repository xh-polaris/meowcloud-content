package redis

import (
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/config"
)

func NewRedis(config *config.Config) *redis.Redis {
	return redis.MustNewRedis(*config.Redis)
}
