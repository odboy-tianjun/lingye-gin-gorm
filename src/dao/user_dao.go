package dao

import (
	"lingye-gin/src/config"
	"lingye-gin/src/entity"
	"lingye-gin/src/service/query"
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
func (UserDAO) SelectPage(condition query.UserQuery) ([]entity.User, int) {
	var users []entity.User
	var total int

	page, pageSize := util.FixPage(condition.Page, condition.PageSize)
	// 分页条件
	pageConfig := config.SqlExcutor.Where("username LIKE ?", "%"+condition.Username+"%")
	// 分页参数
	pageConfig = pageConfig.Limit((page - 1) * pageSize).Offset(pageSize)
	pageConfig = pageConfig.Find(&users)
	// 统计数量
	pageConfig.Count(&total)
	return users, total
}

func (UserDAO) UpdateAll(user *entity.User) {
	config.SqlExcutor.Save(&user)
}

// 更新特定字段
func (UserDAO) UpdateById(user *entity.User) {
	model := &entity.User{}
	model.ID = user.ID
	config.SqlExcutor.Model(model).Updates(map[string]interface{}{
		"username": user.UserName,
	})
}

// 删除
func (UserDAO) DeleteById(id uint) {
	user := entity.User{}
	user.ID = id
	config.SqlExcutor.Delete(&user)
}
