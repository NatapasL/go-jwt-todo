package persistences

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/NatapasL/go-jwt-todo/repositories"
)

type redisAuthTokenRepository struct {
	Redis *redis.Client
}

func NewRedisAuthTokenRepository(r *redis.Client) repositories.AuthTokenRepository {
	return &redisAuthTokenRepository{Redis: r}
}

func (r redisAuthTokenRepository) Find(key string) (string, error) {
	result, err := r.Redis.Get(key).Result()
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (r redisAuthTokenRepository) Create(key string, value string, expiration time.Duration) error {
	err := r.Redis.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r redisAuthTokenRepository) Delete(key string) error {
	_, err := r.Redis.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}
