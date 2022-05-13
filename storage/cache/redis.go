package cache

import "github.com/go-redis/redis"

type RedisFatherInterface interface {
	GetAllKeys() []string
}

type ListInterface interface {
	RedisFatherInterface
	GetNoWait(key string) (string, error) // ~di: interface{}类型、数据压缩
	Get(key string, timeout int) (string, error)
	Put(key string, value string, timeout int) (int, error)
	PutNoWait(key string, value string) (int, error)
	QSize(key string) int
	Empty(key string) bool
	Full(key string) bool
}

type RedisInterface interface {
	ListInterface
}

type Res struct {
	Client *redis.Client
}

type Redis struct {
	Redis RedisConfig `json:"Redis" yaml:"redis"`
}

type RedisConfig struct {
	Network  string `json:"Network" yaml:"network"`
	Addr     string `json:"Addr" yaml:"addr"`
	Password string `json:"Password" yaml:"password"`
	DB       int    `json:"DB" yaml:"db"`
}

func (Res) NewRedisClient(Redis Redis) (*redis.Client, error) {

	options := &redis.Options{
		Network:  Redis.Redis.Network,
		Addr:     Redis.Redis.Addr,
		Password: Redis.Redis.Password,
		DB:       Redis.Redis.DB,
	}

	newClient := redis.NewClient(options)

	return newClient, nil
}
