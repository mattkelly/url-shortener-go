package db

// TODO make db an interface instead and have a redis db implementation

import (
	"os"
	"strings"

	"github.com/go-redis/redis"
)

var client *redis.Client

func getRedisAddr() string {
	redis_host := os.Getenv("REDIS_HOST")
	if redis_host == "" {
		redis_host = "localhost"
	}

	redis_port := os.Getenv("REDIS_PORT")
	if redis_port == "" {
		redis_port = "6379"
	}

	return strings.Join([]string{redis_host, redis_port}, ":")
}

func getRedisPassword() string {
	// Default is ""
	return os.Getenv("REDIS_PASSWORD")
}

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     getRedisAddr(),
		Password: getRedisPassword(),
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
