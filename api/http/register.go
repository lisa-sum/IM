package http

import (
	"context"
	"fmt"
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
	var userInfo UserBasic
	if err := c.ShouldBind(&userInfo); err != nil {
		log.Println("转换JSON失败")
		return
	}

	account := c.PostForm("account")
	fmt.Println("account:", account)

	result, err := db.Mongo.
		Database("im").
		Collection(schema.UserBasic{}.Collection()).
		InsertOne(context.TODO(), userInfo)
	log.Println("result:", result)
	if err != nil {
		log.Println("插入数据失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    "OK",
		"message": "注册成功",
		"code":    http.StatusOK,
	})
}
