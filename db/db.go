package db

// TODO make db an interface instead and have a redis db implementation

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() {
	// TODO accept redis config
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Set(key, value string) error {
	return client.Set(key, value, 0).Err()
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}

func Exists(key string) bool {
	return client.Exists().Val() != 0
}
