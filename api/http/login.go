package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"im/db"
	"im/schema"
	"log"
	"net/http"
)

type DatabaseBasic struct {
	Account  string `bson:"account" json:"account"`
	Password string `bson:"password"`
}

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	log.Println("account", account)
	log.Println("password", password)

	var user schema.UserBasic
	fmt.Println(user.Account)
	var dbBasic DatabaseBasic

	collection := db.Mongo.Database("im").Collection("user_basic")
	if account != "" || password != "" {
		filter := bson.D{{"account", account}}
		err := collection.FindOne(context.TODO(), filter).Decode(&dbBasic)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"body":    "该用户未注册",
					"message": "数据库不存在该账号",
					"code":    http.StatusForbidden,
				})
				return
			}
			log.Println(err)
			return
		}
	} else if account != dbBasic.Password && password != dbBasic.Password {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"body":    "账号密码错误",
			"message": "账号密码错误",
			"code":    http.StatusUnauthorized,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"body":    "OK",
		"message": "登陆成功",
		"code":    http.StatusOK,
	})
}
