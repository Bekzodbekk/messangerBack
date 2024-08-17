package redis

import (
	"context"
	"fmt"
	"log"
	config "user-service/internal/user/pkg/load"

	"github.com/redis/go-redis/v9"
)

func InitRedis(conf config.Config) (*redis.Client, error) {
	target := fmt.Sprintf("%s:%d", conf.Redis.RedisHost, conf.Redis.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr: target,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		return nil, err
	}

	return rdb, nil
}
