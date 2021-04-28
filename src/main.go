package main

import (
	"lingye-gin/src/config"
	"lingye-gin/src/middleware"
)

func main() {
	// 初始化yaml配置
	new(config.ApplicationProperties).Init()
	// 初始化redis
	new(middleware.RedisPool).Init()
	// 初始化gorm
	new(config.DataSourcePool).Connect().LoadEntity()
	// 延时调用函数
	defer config.SqlExcutor.Close()
	// 加载jwt
	new(middleware.JwtEngine).Init()
	// 初始化gin
	new(middleware.GinEngine).Start()
}
