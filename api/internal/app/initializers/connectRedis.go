package initializers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         context.Context
)

func ConnectRedis(config *Config) {
	ctx = context.TODO()
	opt, err := redis.ParseURL(config.RedisUri)

	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(opt)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	fmt.Println("Redis connected!")
}
