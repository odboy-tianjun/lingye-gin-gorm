package entity

// 用户和角色关联
type UserRole struct {
	UserId uint64 `gorm:"column:user_id;type:bigint(20);not null;comment:'用户ID'"`
	RoleId uint64 `gorm:"column:role_id;type:bigint(20);not null;comment:'角色ID'"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
