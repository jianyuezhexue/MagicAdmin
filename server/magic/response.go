package magic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 返回结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功返回
func Success(c *gin.Context, code int, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}

// Fail 成功返回
func Fail(c *gin.Context, code int, msg string, data ...interface{}) {
	// data如何默认
	c.JSON(code, gin.H{"code": code, "msg": msg, "data": data})
}
