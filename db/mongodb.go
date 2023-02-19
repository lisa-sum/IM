package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Mongo *mongo.Client
	Err   error
)

func MongoDBInit() {
	//filePath := utils.GetFilePath("db.yaml")
	//viper.SetConfigFile(filePath)
	//readErr := viper.ReadInConfig()
	//if readErr != nil {
	//	log.Fatal("读取配置失败!" + readErr.Error())
	//}
	//
	//// 给viper读取文件的路径
	//username := viper.GetString("mongodb.username")
	//password := viper.GetString("mongodb.password")
	//url := viper.GetString("mongodb.url")

	Mongo, Err = mongo.Connect(context.Background(), options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "msdnmm",
	}).ApplyURI("mongodb://192.168.0.152:27017"))
	if Err != nil {
		log.Fatal("运行Mongodb服务失败!" + Err.Error())
	}
}
