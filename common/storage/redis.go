package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var RedisS *redis.Client

func InitRedis() (err error) {

	var (
		result string
	)

	RedisS = redis.NewClient(&redis.Options{
		Addr:     "114.117.198.50:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err = RedisS.Ping(ctx).Result()
	fmt.Println(result)
	return err
}
