package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var Redis *redis.Client

func RedisDBInit() {
	Redis = redis.NewClient(&redis.Options{
		Addr:        "192.168.0.158:6379",
		Password:    "msdnmm",
		DB:          0,
		DialTimeout: 30 * time.Second,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		log.Fatal("Ping Redis Error:" + err.Error())
	}
	//defer func(client *redis.Client) {
	//	err := client.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(Redis)
	fmt.Println("Successfully connected redis and pinged.")
}
