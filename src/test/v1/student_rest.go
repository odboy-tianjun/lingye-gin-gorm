package v1

import (
	"fmt"
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

func DescribeStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id=", id)
	var student Student
	if student.ID == 0 {
		c.JSON(404, gin.H{"message": "student not found"})
		return
	}
	c.JSON(200, student)
}

func CreateStudent(c *gin.Context) {
	var student Student
	// 绑定一个请求主体到一个类型
	_ = c.BindJSON(&student)
	c.JSON(200, "创建成功")
}

func ModifyStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id=", id)
	var student Student
	if student.ID == 0 {
		c.JSON(404, gin.H{"message": "student not found"})
		return
	} else {
		_ = c.BindJSON(&student)
		c.JSON(200, student)
	}
}

func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id=", id)
	var student Student
	if student.ID == 0 {
		c.JSON(404, gin.H{"message": "student not found"})
		return
	} else {
		_ = c.BindJSON(&student)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}
