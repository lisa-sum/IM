package service

import (
	"github.com/gin-gonic/gin"
	api "im/api/http"
	"im/schema"
	"net/http"
)

func Login(c *gin.Context) {
	var user = new(schema.UserBasic)
	account := c.PostForm("account")
	password := c.PostForm("password")

	api.Login(account, password, user)

	c.JSON(http.StatusOK, schema.Status{
		Code:    http.StatusOK,
		Message: "登陆成功",
		Body:    user,
	})
}
