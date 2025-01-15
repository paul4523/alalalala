package repo

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type Casher struct {
	Cash redis.Client
}

func NewCash() Casher {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	return Casher{
		Cash: *redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		}),
	}
}
func (c Casher) Set(key string, value interface{}) error {
	return c.Cash.Set(key, value, 0).Err()
}

func (c Casher) Get(key string) (string, error) {
	return c.Cash.Get(key).Result()
}
