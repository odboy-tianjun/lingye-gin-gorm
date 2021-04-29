package v2

import (
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
