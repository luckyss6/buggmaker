package cache

import "github.com/go-redis/redis"

type Res struct {
	Client *redis.Client
}

type RedisConfig struct {
	Network  string `json:"Network" yaml:"network"`
	Addr     string `json:"Addr" yaml:"addr"`
	Password string `json:"Password" yaml:"password"`
	DB       int    `json:"DB" yaml:"db"`
}

func NewRedisClient(redisConfig RedisConfig) (*redis.Client, error) {

	options := &redis.Options{
		Network:  redisConfig.Network,
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	}

	newClient := redis.NewClient(options)

	return newClient, nil
}
