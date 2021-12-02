package redis

import (
	"github.com/go-redis/redis"
	"time"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB: 0,
})

func SetValue(key string, value string) error {
	err := redisClient.Set(key, value, 1 * time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetValue(key string) (string, error) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}

	return value,nil
}
