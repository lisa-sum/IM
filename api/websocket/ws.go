package websocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im/db"
	"im/schema"
	"log"
	"net/http"
)

type MessageBasic2 struct {
	UserIdentity string `json:"user_identity" bson:"user_identity"`
	RoomIdentity string `json:"room_identity" bson:"room_identity"`
	Data         string `json:"data" bson:"data"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
	UpdatedAt    string `json:"updated_at" bson:"updated_at"`
}

// 处理WS跨域
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var ws = make(map[string]*websocket.Conn)

func Ws(c *gin.Context) {
	// 升级为WS协议 s
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	// 升级为WS协议 e

	for {
		// 创建消息JSON结构体, 保存消息与额外的信息
		message := new(schema.MessageBasic)

		// ReadJSON
		// message: 需要读取的消息对象(一条消息一般包含多个属性用于其他用途)
		err := conn.ReadJSON(message)
		if err != nil {
			log.Println("read:", err)
			break
		}

		// 消息结构体
		msg := schema.MessageBasic{
			UserIdentity: message.UserIdentity,
			RoomIdentity: message.RoomIdentity,
			Data:         message.Data,
			CreatedAt:    message.CreatedAt,
			UpdatedAt:    message.UpdatedAt,
		}

		// 存储消息
		go saveMessage(c, msg)

		// 根据连接绑定用户id
		ws[msg.UserIdentity] = conn
		for _, cc := range ws {
			// WriteMessage
			// 1 消息类型: websocket.TextMessage文本
			// 2 传输类型: []byte二进制
			err = cc.WriteJSON(msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func saveMessage(c *gin.Context, msg schema.MessageBasic) {

	// 转换JSON s
	jsonMessage, err := json.Marshal(&msg)
	if err != nil {
		log.Println("转换JSON失败:" + err.Error())
	}

	// 存储消息, 存储异常则返回消息给客户端,继续执行
	if err := db.Redis.RPush(msg.UserIdentity, jsonMessage).Err(); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"body":    nil,
			"message": "存储消息失败",
			"code":    http.StatusBadRequest,
		})
		log.Println("存储消息失败:" + err.Error())
	}
}
