package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err == nil {
		log.Fatalln("存储文件失败" + err.Error())
	}
	log.Println(file.Filename)
	//c.JSON(http.StatusOK, gin.H{
	//	"code":    http.StatusOK,
	//	"body":    file.Filename,
	//	"message": "存储文件成功",
	//})

	//uploadDir := utils.GetDirPath("uploads/")
	//uploadDir := path.Dir("E:\\Web\\Golang\\src\\02\\11\\im\\uploads")
	//log.Println("uploadDir:", uploadDir)
	////utils.Log(uploadDir)
	////uploadInfo, err := api.UploadFileToMinio(file.Filename, uploadDir)
	//// 读取minio配置文件字段
	//url := viper.GetString("192.168.0.158:5000")
	//accessKey := viper.GetString("3YSfBdh4C6XVMgV4")
	//secretKey := viper.GetString("m2nClKEI823yY6ffC6DjdgXQTpNf2m6h")
	//
	//// 初使化minio client对象。
	//var err error
	//var MinioClient *minio.Client
	//MinioClient, err = minio.New(url, &minio.Options{
	//	Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
	//	Secure: false,
	//})
	//if err != nil {
	//	log.Fatal("初始化Minio客户端失败" + err.Error())
	//}
	//
	//// 上传文件至minio
	//log.Println("file", file)
	//bucketName := "chat"
	//
	//uploadInfo, err := MinioClient.FPutObject(
	//	context.Background(),
	//	bucketName,
	//	file.Filename,
	//	uploadDir,
	//	minio.PutObjectOptions{
	//		ContentType: "application/svg",
	//	})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println("Successfully uploaded object: ", uploadInfo)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, gin.H{
	//		"message": "上传文件至文件存储服务器失败:" + err.Error(),
	//		"body":    nil,
	//		"code":    http.StatusInternalServerError,
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "上传文件至文件存储服务器失败",
	//	"body":    "uploadInfo",
	//	"code":    http.StatusOK,
	//})
}
