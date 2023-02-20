package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ws "im/api/websocket"
	"im/schema"
	"net/http"
)

// 处理WebSocket跨域
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var wc = make(map[string]*websocket.Conn)

func IM(c *gin.Context) {
	// 升级为WebSocket协议 s
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    http.StatusInternalServerError,
			Message: "系统异常：" + err.Error(),
			Body:    nil,
		})
		return
	}

	ws.Chat(c, conn, wc)
}
