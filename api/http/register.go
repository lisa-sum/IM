package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"im/db"
	"im/schema"
	"log"
	"net/http"
)

type UserBasic struct {
	Account   string `bson:"account" form:"account"`
	Password  string `bson:"password" form:"password"`
	Avatar    string `bson:"avatar" form:"avatar"`
	Email     string `bson:"email" form:"email"`
	Nickname  string `bson:"nickname" form:"nickname"`
	Gender    string `bson:"gender" form:"gender"`
	CreatedAt int    `bson:"created_at" form:"createdAt"`
	UpdatedAt int    `bson:"updated_at" form:"updatedAt"`
}

func Register(c *gin.Context) {
	// 创建用户结构体,用于绑定JSON数据
	var userInfo UserBasic

	// JSON转为结构体,转换失败抛出异常给客户端与输出异常
	if err := c.ShouldBind(&userInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":    "转换JSON失败",
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		log.Println("转换JSON失败")
		return
	}

	// 创建用户
	result, err := db.Mongo.
		Database("im").
		Collection(schema.UserBasic{}.Collection()).
		InsertOne(context.TODO(), userInfo)
	log.Println("result:", result)
	// 创建用户时的异常处理
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":    http.StatusBadRequest,
			"message": "创建用户异常",
			"code":    http.StatusOK,
		})
		log.Println("创建用户异常" + err.Error())
		return
	}
	// 成功返回客户端创建好的数据条目的ID
	c.JSON(http.StatusOK, gin.H{
		"data":    result.InsertedID,
		"message": "注册成功",
		"code":    http.StatusOK,
	})
}
