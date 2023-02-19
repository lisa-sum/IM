package test

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"testing"
	"time"
)

func TestRedisDB(t *testing.T) {
	configure := GetFilePath("db.yaml")
	viper.SetConfigFile(configure)
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
	url := viper.GetString("redis.url")
	password := viper.GetString("redis.password")

	client := redis.NewClient(&redis.Options{
		Addr:        url,
		Password:    password,
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

	t.Log("Successfully connected redis and pinged." + pong)
}

func GetFilePath(args ...string) string {
	var dir string
	if len(args) < 2 || args[1] == "" {
		dir = "/config/"
	}
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))    //获取当前工作目录
	dirPath := path.Dir(root + dir)         // 获取配置文件的目录
	filePath := path.Join(dirPath, args[0]) // 获取配置文件
	return filePath
}
