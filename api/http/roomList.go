package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"im/db"
	"im/schema"
	"log"
	"net/http"
	"strconv"
)

func RoomList(c *gin.Context) {
	query := c.Query("number")
	dbBasic := db.Mongo.Database("im").Collection(schema.RoomBasic{}.Collection())
	if query == "all" { // 查询全部群聊
		// 查询全部
		filter := bson.D{{}}
		// 只返回name字段
		opts := options.Find().SetProjection(bson.D{
			{"info", 0},
			{"user_identity", 0},
			{"created_at", 0},
		})
		cursor, err := dbBasic.Find(context.TODO(), filter, opts)
		var roomList = make([]map[string]any, 0)
		if err = cursor.All(context.TODO(), &roomList); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"body":    nil,
				"code":    http.StatusNotAcceptable,
				"message": "请求所有群聊失败",
			})
			log.Println("请求所有群聊失败" + err.Error())
			return
		}
		log.Println("roomList:", roomList)
		c.JSON(http.StatusOK, gin.H{
			"body":    roomList,
			"code":    http.StatusOK,
			"message": "请求所有群聊成功",
		})
		return
	} else if _, err := strconv.Atoi(query); err != nil { // 如果转为int类型成功,那么为群号直接查询
		// 根据传入的number查询
		var roomInfo schema.RoomBasic
		filter := bson.D{{"number", query}}
		err = dbBasic.FindOne(context.TODO(), filter).Decode(&roomInfo)
		log.Println("roomInfo:", roomInfo)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"body":    nil,
					"code":    http.StatusNotAcceptable,
					"message": "查询的群号无结果",
				})
				return
			}
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"body":    roomInfo,
			"code":    http.StatusOK,
			"message": fmt.Sprintf("获取number为%v的结果成功", query),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"body":    nil,
		"code":    http.StatusBadRequest,
		"message": "参数异常, 请传递正确的参数",
	})
	log.Println("异常/或未携带值的请求参数")
}
