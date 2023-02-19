package test

import (
	"github.com/go-redis/redis"
	"os"
	"testing"
	"time"
)

func TestRedisDB(t *testing.T) {
	var Redis *redis.Client
	Redis = redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          0,
		DialTimeout: 30 * time.Second,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		t.Fatal("Ping Redis Error:" + err.Error())
	}
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(Redis)
	t.Log("Successfully connected redis and " + pong)
}
