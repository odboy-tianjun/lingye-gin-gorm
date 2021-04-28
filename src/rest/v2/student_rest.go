package v2

import (
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
}

func DescribeStudents(c *gin.Context) {
	var students []Student
	c.JSON(200, gin.H{
		"data":     students,
		"page":     1,
		"pageSize": 15,
	})
}
