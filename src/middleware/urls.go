package middleware

import (
	"github.com/gin-gonic/gin"
	"lingye-gin/src/modules/system/rest"
	"lingye-gin/src/test"
	v1 "lingye-gin/src/test/v1"
	v2 "lingye-gin/src/test/v2"
	"lingye-gin/src/util"
)

type RequestApi struct {
	// get、post、delete、put
	Mode string
	// 分组名称
	GroupName string
	// 分组过滤器
	GroupHandleFunction gin.HandlerFunc
	// 请求路径
	RelativePath string
	// 请求处理器
	HandleFunction gin.HandlerFunc
}

// 定义变长数组变量
var Urls = [...]RequestApi{
	// 定义请求方式和路径
	{Mode: "get", RelativePath: "/sn", HandleFunction: test.DescribeSign},
	// 获取所有用户
	{Mode: "get", RelativePath: "/users", HandleFunction: rest.DescribeUsers},
	// v1
	// 获取所有学生
	{GroupName: "v1", Mode: "get", RelativePath: "/students", HandleFunction: v1.DescribeStudents},
	// 根据ID获取学生
	{GroupName: "v1", Mode: "get", RelativePath: "/students/:id", HandleFunction: v1.DescribeStudentById},
	// 保存学生
	{GroupName: "v1", Mode: "post", RelativePath: "/students", HandleFunction: v1.CreateStudent},
	// 根据ID更新学生
	{GroupName: "v1", Mode: "put", RelativePath: "/students/:id", HandleFunction: v1.ModifyStudentById},
	// 根据ID删除学生
	{GroupName: "v1", Mode: "delete", RelativePath: "/students/:id", HandleFunction: v1.DeleteStudentById},
	// v2
	{GroupName: "v2", GroupHandleFunction: util.VerifySign, Mode: "get", RelativePath: "/students", HandleFunction: v2.DescribeStudents},
	// api jwt
	// 获取所有用户
	{GroupName: "api", Mode: "get", RelativePath: "/users", HandleFunction: rest.DescribeUsers},
	// 根据ID获取用户
	{GroupName: "api", Mode: "get", RelativePath: "/users/:id", HandleFunction: rest.DescribeUserById},
	// 保存用户
	{GroupName: "api", Mode: "post", RelativePath: "/users", HandleFunction: rest.CreateUser},
	// 根据ID更新用户
	{GroupName: "api", Mode: "put", RelativePath: "/users/:id", HandleFunction: rest.ModifyUserById},
	// 根据ID删除用户
	{GroupName: "api", Mode: "delete", RelativePath: "/users/:id", HandleFunction: rest.DeleteUserById},
}
