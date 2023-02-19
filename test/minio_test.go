package test

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"testing"
)

var MinioClient *minio.Client

func TestMinioNewClient(t *testing.T) {
	// 读取minio配置文件
	configure := GetFilePath("minio.yaml")
	t.Log(configure)
	viper.SetConfigFile(configure)
	readErr := viper.ReadInConfig()
	if readErr != nil {
		t.Fatal("读取配置失败!" + readErr.Error())
		return
	}

	// 读取minio配置文件字段
	url := viper.GetString("minio.url")
	accessKey := viper.GetString("minio.accessKey")
	secretKey := viper.GetString("minio.secretKey")

	// 初使化minio client对象。
	var err error
	MinioClient, err = minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		t.Fatal("初始化Minio客户端失败" + err.Error())
	}
	t.Log("测试初始化Minio客户端成功", MinioClient)
}
