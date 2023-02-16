package router

import (
	"github.com/gin-gonic/gin"
	api "im/api/http"
	ws "im/api/websocket"
	"im/utils"
	"log"
)

func Server() {
	server := gin.Default()
	server.Use(utils.Cors())
	//server.GET("/historyMessage", api.MessageList)
	server.PUT("/register", api.Register)
	server.POST("/login", api.Login)
	server.GET("/chat", ws.Ws)
	server.GET("/roomList", api.RoomList)
	server.GET("/historyMessage", api.HistoryMessage)
	err := server.Run(":4000")
	if err != nil {
		log.Println("运行gin服务失败,请检查端口是否被占用", err.Error())
	}
}
