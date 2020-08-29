package infra

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	if client != nil {
		return client
	}

	if err := initRedisClient(); err != nil {
		panic(err)
		return nil
	}

	return client
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
