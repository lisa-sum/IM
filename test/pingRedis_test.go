package test

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestRedisDB(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:        "192.168.0.158:6379",
		Password:    "msdnmm",
		DB:          0,
		DialTimeout: 30 * time.Second,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		t.Fatal("Ping Redis Error:" + err.Error())
	}
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(client)

	fmt.Println("Successfully connected redis and pinged." + pong)
}
