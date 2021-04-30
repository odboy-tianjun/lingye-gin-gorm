package main

import (
	"fmt"
	"lingye-gin/src/config"
	"lingye-gin/src/middleware"
	"lingye-gin/src/modules/system/entity"
	"lingye-gin/src/util"
)

func main() {
	// 初始化yaml配置
	new(config.ApplicationProperties).Init()
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

func Test() {
	// 通过ID查询
	var role entity.Role
	util.SelectOne(role, "select * from sys_role where id = ?", 1)
	fmt.Printf("SelectOne=%v", role)
	// 查询列表
}
