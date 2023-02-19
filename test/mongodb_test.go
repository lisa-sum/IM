package test

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestMongoDB(t *testing.T) {
	// 读取mongodb路径
	filePath := GetFilePath("db.yaml")
	fmt.Println("filePath:", filePath)
	viper.SetConfigFile(filePath)
	readErr := viper.ReadInConfig()
	if readErr != nil {
		log.Fatal("读取配置失败!" + readErr.Error())
	}

	username := viper.GetString("mongodb.username") // 解析配置的属性
	password := viper.GetString("mongodb.password") // 解析配置的属性
	uri := viper.GetString("mongodb.url")           // 解析配置的属性

	client, err := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: username,
		Password: password,
	}).ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			t.Fatal(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}
	fmt.Println("Successfully connected mongodb and pinged.")
}
