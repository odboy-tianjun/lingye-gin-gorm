package dao

import (
	"lingye-gin/src/config"
	"lingye-gin/src/entity"
	"lingye-gin/src/service/dto"
	"lingye-gin/src/util"
)

type UserDAO struct{}

// 新增
func (UserDAO) Insert(resource *entity.User) {
	config.SqlExcutor.Create(resource)
}

// ID查询
func (UserDAO) SelectOne(id uint) entity.User {
	var user entity.User
	config.SqlExcutor.First(&user, id)
	return user
}

// 查询所有记录
func (UserDAO) SelectAll() []entity.User {
	var users []entity.User
	config.SqlExcutor.Find(&users)
	return users
}

// 条件查询
func (UserDAO) SelectList(resource entity.User) []entity.User {
	var users []entity.User
	config.SqlExcutor.Where("username LIKE ?", "%"+resource.UserName+"%").Find(&users)
	return users
}

// 条件分页查询
func (UserDAO) SelectPage(dto dto.UserDTO) []entity.User {
	var users []entity.User

	page, pageSize := util.FixPage(dto.Page, dto.PageSize)
	// 分页条件
	pageConfig := config.SqlExcutor.Where("username LIKE ?", "%"+dto.Username+"%")
	// 分页参数
	pageConfig = pageConfig.Limit((page - 1) * pageSize).Offset(pageSize)
	pageConfig.Find(&users)
	return users
}

func test() {
	//
	//user := entity.User{Username: "lingye"}
	//
	//userDAO := dao.UserDAO{}
	//userDAO.Insert(&user)
	//userDAO.SelectAll()
	//userDAO.SelectOne(1)
	//userDAO.SelectList(user)
	//
	//dto := dto2.UserDTO{}
	//dto.Username = "测试"
	//userDAO.SelectPage(dto)
}
