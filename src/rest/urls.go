package rest

import (
	"github.com/gin-gonic/gin"
	v3 "lingye-gin/src/rest/api"
	v1 "lingye-gin/src/rest/v1"
	v2 "lingye-gin/src/rest/v2"
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
	{Mode: "get", RelativePath: "/sn", HandleFunction: DescribeSign},
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
	{GroupName: "api", Mode: "get", RelativePath: "/students", HandleFunction: v3.DescribeStudents},
}
