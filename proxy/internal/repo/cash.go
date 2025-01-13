package repo

import "github.com/go-redis/redis"

type Casher struct {
	Cash redis.Client
}

func NewCash() Casher {
	return Casher{
		Cash: *redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
}
func (c Casher) Set(key string, value interface{}) error {
	return c.Cash.Set(key, value, 0).Err()
}

func (c Casher) Get(key string) (string, error) {
	return c.Cash.Get(key).Result()
}
