package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"im/db"
	"im/schema"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	log.Println("account", account)
	log.Println("password", password)

	var user schema.UserBasic

	collection := db.Mongo.Database("im").Collection("user_basic")
	if account != "" || password != "" {
		filter := bson.D{{"account", account}}
		err := collection.FindOne(context.TODO(), filter).Decode(&user)
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
	} else if account != user.Password && password != user.Password {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"body":    "账号密码错误",
			"message": "账号密码错误",
			"code":    http.StatusUnauthorized,
		})
		return
	}
	log.Println("user:", user)
	c.JSON(http.StatusOK, gin.H{
		"body":    user,
		"message": "登陆成功",
		"code":    http.StatusOK,
	})
}
