package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Body    any `json:"body"`
}

var res Status

func send(ok bool, args ...interface{}) {
	fmt.Print("ok:", ok)
	if !ok {
		res.Message = args[1]
		res.Body = args[2]
		if len(args) <= 3 {
			res.Code = http.StatusBadRequest
		} else {
			res.Code = args[3].(int)
		}
		args[0].(*gin.Context).AbortWithStatusJSON(http.StatusOK, res)
		return
	}
	res.Message = args[1]
	res.Body = args[2]
	if len(args) <= 3 {
		res.Code = http.StatusOK
	} else {
		res.Code = args[3].(int)
	}
	args[0].(*gin.Context).JSON(http.StatusOK, res)
}

// SendOk 返回返回自定义正确的状态体
/**
 * @description 返回的数据体类型规范化
 * @since 23/01/20232:35 pm
 * @param args[0] *gin.Context
 * @param args[1] Message 返回消息
 * @param args[2] Body 返回内容
 * @param args[3] Code 返回状态码
 * @return void
 *  */
func SendOk(args ...interface{}) {
	//send(true, args)
	res.Message = args[1]
	res.Body = args[2]
	if len(args) <= 3 {
		res.Code = http.StatusOK
	} else {
		res.Code = args[3].(int)
	}
	args[0].(*gin.Context).JSON(http.StatusOK, res)
}

// SendBad 返回自定义错误的状态体
/**
 * @description 返回的数据体类型规范化
 * @since 23/01/20232:35 pm
 * @param args[0] *gin.Context
 * @param args[1] Message 返回消息
 * @param args[2] Body 返回内容
 * @param args[3] Code 返回状态码
 * @return void
 *  */
func SendBad(args ...interface{}) {
	//send(false, args)
	res.Message = args[1]
	res.Body = args[2]
	if len(args) <= 3 {
		res.Code = http.StatusBadRequest
	} else {
		res.Code = args[3].(int)
	}
	args[0].(*gin.Context).AbortWithStatusJSON(http.StatusOK, res)
}
