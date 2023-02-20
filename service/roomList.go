package service

import (
	"github.com/gin-gonic/gin"
	api "im/api/http"
	"log"
	"net/http"
)

func RoomList(c *gin.Context) {
	query := c.Query("number")

	result, err := api.RoomList(query)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, result)
}
