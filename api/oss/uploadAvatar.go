package oss

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"im/db"
	"im/schema"
	"log"
	"os"
	"path/filepath"
)

func UploadAvatar(args ...any) *mongo.UpdateResult {
	go UploadFile(args[0].(string))

	// 协议
	protocol := "http://"
	// 文件存储服务器地址+port
	host := os.Getenv("MINIO_URL")
	// 存储桶名
	bucket := os.Getenv("MINIO_BUCKET")
	// 转为Linux文件路径
	filepathToLinux := filepath.ToSlash(filepath.Join(host, bucket, args[0].(string)))
	// 完整URI
	uri := protocol + filepathToLinux

	// 对传入的ObjectID文档进行修改, 添加传入的头像URI
	filter := bson.D{{"_id", args[1]}}
	update := bson.D{{"$set", bson.D{{"avatar", uri}}}}
	result, err := db.Mongo.Database("im").
		Collection(schema.UserBasic{}.Collection()).
		UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result
}
