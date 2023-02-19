package test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestMongoDB(t *testing.T) {
	ENV, err := godotenv.Read()
	if err != nil {
		log.Fatal("获取环境变量失败:", err.Error())
	}

	client, err := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{
		Username: ENV["MONGODB_USERNAME"],
		Password: ENV["MONGODB_PASSWORD"],
	}).ApplyURI(ENV["MONGODB_URL"]))
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
