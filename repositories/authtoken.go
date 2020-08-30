package repositories

import (
	"time"
)

type AuthTokenRepository interface {
	Find(key string) (string, error)
	Create(key string, value string, expiration time.Duration) error
	Delete(key string) error
}
