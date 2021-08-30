package initialize

import (
	"fmt"
	"go-web/internal/pkg/config"

	"github.com/go-redis/redis"
)

// 初始化Redis
func Redis(redisConf *config.RedisConfiguration) *redis.Client {

	redisIns := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})

	err := redisIns.Ping().Err()
	if err != nil || redisIns == nil {
		panic(fmt.Sprintf("初始化Redis异常：%v", err))
	}

	return redisIns
}
