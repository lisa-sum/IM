package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	Mongo *mongo.Client
	Err   error
)

func MongoDBInit() {
	Mongo, Err = mongo.Connect(context.Background(), options.Client().SetAuth(options.Credential{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	}).ApplyURI(os.Getenv("MONGODB_URL")))
	if Err != nil {
		log.Fatal("运行Mongodb服务失败!" + Err.Error())
	}
}
