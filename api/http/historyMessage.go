package http

import (
	"github.com/gin-gonic/gin"
	"im/db"
	"log"
	"net/http"
)

func HistoryMessage(c *gin.Context) {
	account := c.Query("account")

	result, err := db.Redis.LRange(account, 0, -1).Result()
	if err != nil {
		log.Println("获取消息列表失败" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"body":    result,
			"message": "获取消息列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"body":    result,
		"message": "获取消息列表成功",
	})
}
