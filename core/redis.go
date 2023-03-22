package core

import (
	"github.com/go-redis/redis"
)

func (this *Common) Redischeck(redishost, redispassword string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     redishost,
		Password: redispassword, // no password set
		DB:       0,             // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
