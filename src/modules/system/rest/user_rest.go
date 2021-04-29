package rest

import (
	"github.com/gin-gonic/gin"
	"lingye-gin/src/modules/system/service"
	"lingye-gin/src/modules/system/service/dto"
	"lingye-gin/src/modules/system/service/query"
	"lingye-gin/src/util"
	"net/http"
)

// 统一Service对象
var userService = &service.UserService{}

func DescribeUsers(c *gin.Context) {
	var userQuery query.UserQuery
	_ = c.BindJSON(&userQuery)

	users, total := userService.DescribeUsers(userQuery)
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  users,
		"total": total,
	})
}

func DescribeUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	user := userService.DescribeUserById(util.StringToUInt(id))
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
	userService.Save(&userDTO)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}

func ModifyUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	localUser := userService.DescribeUserById(util.StringToUInt(id))
	if localUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "user not found",
		})
		return
	} else {
		var userDTO dto.UserDTO
		_ = c.BindJSON(&userDTO)
		userService.ModifyById(&userDTO)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
		})
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	uid := util.StringToUInt(id)
	localUser := userService.DescribeUserById(uid)
	if localUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "user not found",
		})
		return
	} else {
		userService.RemoveById(uid)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
		})
	}
}
