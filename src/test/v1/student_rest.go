package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lingye-gin/src/base"
	"lingye-gin/src/util"
)

type Student struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
}

func DescribeStudents(c *gin.Context) {
	var students []Student
	var pageData base.SimplePageData

	pageData.Total = 0
	pageData.Data = students
	util.RSuccessWithMsgJson(c, "分页获取学生列表成功!", pageData)
}

func DescribeStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id=", id)
	var student Student
	if student.ID == 0 {
		util.RErrorJson(c, "没有找到该学生信息", nil)
		return
	}
	util.RSuccessJson(c, student)
}

func CreateStudent(c *gin.Context) {
	var student Student
	// 绑定一个请求主体到一个类型
	_ = c.BindJSON(&student)
	util.RSuccessWithMsgJson(c, "创建成功", nil)
}

func ModifyStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id=", id)
	var student Student
	if student.ID == 0 {
		util.RErrorJson(c, "非法操作, 没有找到该学生信息!", nil)
		return
	} else {
		_ = c.BindJSON(&student)
		util.RSuccessJson(c, student)
	}
}

func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id=", id)
	var student Student
	if student.ID == 0 {
		util.RErrorJson(c, "非法操作, 没有找到该学生信息!", nil)
		return
	} else {
		_ = c.BindJSON(&student)
		util.RSuccessMsg(c, "删除成功")
	}
}
