package infra

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func GetRedisClient() (*redis.Client, error) {
	if client != nil {
		return client, nil
	}

	if err := initRedisClient(); err != nil {
		return nil, err
	}

	return client, nil
}

func initRedisClient() error {
	addr := os.Getenv("REDIS_URL")
	if len(addr) <= 0 {
		addr = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
