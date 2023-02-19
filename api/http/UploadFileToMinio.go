package http

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
)

var MinioClient *minio.Client

func UploadFileToMinio(file string, dir string) (any, error) {
	//// 读取minio配置文件
	//configure := utils.GetFilePath("minio.yaml")
	//fmt.Print(configure)
	//viper.SetConfigFile(configure)
	//readErr := viper.ReadInConfig()
	//if readErr != nil {
	//	log.Fatal("读取配置失败!" + readErr.Error())
	//	return "", readErr
	//}

	// 读取minio配置文件字段
	url := viper.GetString("192.168.0.158:5000")
	accessKey := viper.GetString("3YSfBdh4C6XVMgV4")
	secretKey := viper.GetString("m2nClKEI823yY6ffC6DjdgXQTpNf2m6h")

	// 初使化minio client对象。
	var err error
	MinioClient, err = minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal("初始化Minio客户端失败" + err.Error())
	}

	// 上传文件至minio
	log.Println("filename:", file)
	bucketName := "chat"
	objectName := file
	filePath := dir

	uploadInfo, err := MinioClient.FPutObject(
		context.Background(),
		bucketName,
		objectName,
		filePath,
		minio.PutObjectOptions{
			ContentType: "application/svg",
		})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)

	return uploadInfo, err
}
