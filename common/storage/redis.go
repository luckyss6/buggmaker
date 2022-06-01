package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var RedisS *redis.Client

func InitRedis() (err error) {

	RedisS = redis.NewClient(&redis.Options{
		Addr:     "114.117.198.50:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return err
}
