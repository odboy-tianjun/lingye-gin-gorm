package dao

import (
	"lingye-gin/src/config"
	"lingye-gin/src/modules/system/entity"
	"lingye-gin/src/modules/system/service/query"
	"lingye-gin/src/util"
)

type RoleDAO struct{}

// 新增
func (dao RoleDAO) Insert(resource *entity.Role) {
	config.SqlExcutor.Create(resource)
}

// 删除
func (RoleDAO) DeleteById(id uint64) {
	role := entity.Role{}
	role.ID = id
	config.SqlExcutor.Delete(&role)
}

// 更新
func (RoleDAO) UpdateById(role *entity.Role) {
	model := &entity.Role{}
	model.ID = role.ID
	config.SqlExcutor.Model(model).Updates(map[string]interface{}{
		"role_key":  role.RoleKey,
		"role_name": role.RoleName,
		"role_sort": role.RoleSort,
	})
}

// 条件分页查询
func (RoleDAO) SelectPage(condition query.UserQuery) ([]entity.Role, int) {
	var users []entity.Role
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
