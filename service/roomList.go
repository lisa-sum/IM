package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	api "im/api/http"
	"im/schema"
	"log"
	"net/http"
)

func RoomList(c *gin.Context) {
	query := c.Query("number")

	result, err := api.RoomList(query)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, schema.Status{
		Body:    result,
		Code:    http.StatusOK,
		Message: fmt.Sprintf("获取number为%v的结果成功", query),
	})
	log.Println("异常/或未携带值的请求参数")
}
