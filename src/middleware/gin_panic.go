package middleware

import (
	"github.com/gin-gonic/gin"
	"lingye-gin/src/config"
	"net/http"
	"runtime/debug"
)

func GinPanic(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			config.Logger.Printf("panic: %v\n", r)
			debug.PrintStack()
			// 封装通用json返回
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	// 加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
