package http

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"im/api/oss"
	"im/db"
	"im/schema"
	"log"
	"mime/multipart"
)

func Register(userInfo schema.UserFormBasic, file *multipart.FileHeader) (any, error) {

	// 创建用户
	result, err := db.Mongo.
		Database("im").
		Collection(schema.UserBasic{}.Collection()).
		InsertOne(context.TODO(), userInfo)
	log.Println("result:", result)
	// 创建用户时的异常处理
	if err != nil {
		return nil, err
	}

	go oss.UploadAvatar(file.Filename, result.InsertedID.(primitive.ObjectID))
	return result.InsertedID, err
}
