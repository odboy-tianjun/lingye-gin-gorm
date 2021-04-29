package rest

import (
	"github.com/gin-gonic/gin"
	"lingye-gin/src/modules/system/service"
	"lingye-gin/src/modules/system/service/dto"
	"lingye-gin/src/modules/system/service/query"
	"lingye-gin/src/util"
)

// 统一Service对象
var userService = &service.UserService{}

func DescribeUsers(c *gin.Context) {
	var userQuery query.UserQuery
	_ = c.BindJSON(&userQuery)
	users, total := userService.DescribeUsers(userQuery)
	util.RSuccessJson(c, util.BuildPageData(users, total))
}

func DescribeUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	user := userService.DescribeUserById(util.StringToUInt64(id))
	if user.ID == 0 {
		util.RErrorMsg(c, "该用户信息未查询到")
		return
	}
	util.RSuccessJson(c, user)
}

func CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	// 绑定一个请求主体到一个类型
	_ = c.BindJSON(&userDTO)
	userService.Save(&userDTO)
	util.RSuccessMsg(c, "创建成功!")
}

func ModifyUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	localUser := userService.DescribeUserById(util.StringToUInt64(id))
	if localUser.ID == 0 {
		util.RErrorMsg(c, "非法操作, 该用户信息未查询到!")
		return
	} else {
		var userDTO dto.UserDTO
		_ = c.BindJSON(&userDTO)
		userService.ModifyById(&userDTO)
		util.RSuccessMsg(c, "修改成功!")
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	uid := util.StringToUInt64(id)
	localUser := userService.DescribeUserById(uid)
	if localUser.ID == 0 {
		util.RErrorMsg(c, "非法操作, 该用户信息未查询到!")
		return
	} else {
		userService.RemoveById(uid)
		util.RSuccessMsg(c, "删除成功!")
	}
}
