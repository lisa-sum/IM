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
	//filePath := utils.GetPath("db.yaml", "")
	//fmt.Println("filePath:", filePath)
	//viper.SetConfigFile(filePath)
	//readErr := viper.ReadInConfig()
	//if readErr != nil {
	//	log.Fatal("读取配置失败!" + readErr.Error())
	//} // 给viper读取文件的路径
	//username := viper.GetString("mongodb.username") // 解析配置的属性
	//password := viper.GetString("mongodb.password") // 解析配置的属性
	//uri := viper.GetString("mongodb.url")           // 解析配置的属性

	Mongo, Err = mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "msdnmm",
	}).ApplyURI("mongodb://192.168.0.152:27017"))
	if Err != nil {
		log.Fatal("运行Mongodb服务失败!" + Err.Error())
	}
}
