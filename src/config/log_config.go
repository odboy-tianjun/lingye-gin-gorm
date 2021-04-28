package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

// 实例化
var Logger = logrus.New()

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := AppProps.Log.File.Path
	logFileName := AppProps.Log.File.Name
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Sprintf("err: %v", err))
	}

	// 设置输出
	Logger.Out = src
	// 设置日志级别
	if strings.Compare(AppProps.Log.Level, "debug") == 0 {
		Logger.SetLevel(logrus.DebugLevel)
	} else if strings.Compare(AppProps.Log.Level, "info") == 0 {
		Logger.SetLevel(logrus.InfoLevel)
	} else if strings.Compare(AppProps.Log.Level, "warn") == 0 {
		Logger.SetLevel(logrus.WarnLevel)
	} else if strings.Compare(AppProps.Log.Level, "error") == 0 {
		Logger.SetLevel(logrus.ErrorLevel)
	}

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// 新增钩子
	Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		// 这个日期是真的牛皮, yyyy-MM-dd hh:mm:ss它不香吗
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
