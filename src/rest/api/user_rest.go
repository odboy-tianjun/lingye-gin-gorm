package api

import (
	"github.com/gin-gonic/gin"
	"lingye-gin/src/service"
	"lingye-gin/src/service/dto"
	"lingye-gin/src/service/query"
	"lingye-gin/src/util"
	"net/http"
)

func DescribeUsers(c *gin.Context) {
	var userQuery query.UserQuery
	_ = c.BindJSON(&userQuery)

	users, total := new(service.UserService).DescribeUsers(userQuery)
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  users,
		"total": total,
	})
}

func DescribeUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	user := new(service.UserService).DescribeUserById(util.StringToUInt(id))
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    user,
	})
}

func CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	// 绑定一个请求主体到一个类型
	_ = c.BindJSON(&userDTO)
	new(service.UserService).Save(&userDTO)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}

func ModifyUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	localUser := new(service.UserService).DescribeUserById(util.StringToUInt(id))
	if localUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "user not found",
		})
		return
	} else {
		var userDTO dto.UserDTO
		_ = c.BindJSON(&userDTO)
		new(service.UserService).ModifyById(&userDTO)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
		})
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	uid := util.StringToUInt(id)
	localUser := new(service.UserService).DescribeUserById(uid)
	if localUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "user not found",
		})
		return
	} else {
		new(service.UserService).RemoveById(uid)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
		})
	}
}
