package main

import (
	"lingye-gin/src/config"
	"lingye-gin/src/middleware"
	v1 "lingye-gin/src/test/v1"
)

func main() {
	// 初始化yaml配置
	new(config.ApplicationProperties).Init()

	v1.DescribeStudents(nil)

	// 初始化redis
	new(middleware.RedisPool).Init()
	// 初始化gorm, 注册表
	new(config.DataSourcePool).Connect().LoadEntity()
	// 延时调用函数
	defer config.SqlExcutor.Close()
	// 加载jwt
	new(middleware.JwtEngine).Init()
	// 初始化gin
	new(middleware.GinEngine).Start()
}
