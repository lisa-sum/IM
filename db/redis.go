package db

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"im/utils"
	"log"
	"time"
)

var Redis *redis.Client

func RedisDBInit() {
	filePath := utils.GetFilePath("db.yaml")
	viper.SetConfigFile(filePath)
	readErr := viper.ReadInConfig()
	if readErr != nil {
		log.Fatal("读取配置失败!" + readErr.Error())
	}

	// 给viper读取文件的路径
	password := viper.GetString("redis.password")
	url := viper.GetString("redis.url")

	Redis = redis.NewClient(&redis.Options{
		//Addr: "192.168.0.158:6379",
		Addr: url,
		//Password: "msdnmm",
		Password:    password,
		DB:          0,
		DialTimeout: 30 * time.Second,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		log.Fatal("Ping Redis Error:" + err.Error())
	}
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Redis)
	log.Println("Successfully connected redis and " + pong)
}
