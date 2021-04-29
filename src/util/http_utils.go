package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RErrorJson(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

func RSuccessJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
	c.Abort()
}

func RSuccessMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
	c.Abort()
}

func RErrorMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  "success",
	})
	c.Abort()
}

func RSuccessWithMsgJson(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}
