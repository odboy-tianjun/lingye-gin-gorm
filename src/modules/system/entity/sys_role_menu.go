package entity

// 角色和菜单关联
type RoleMenu struct {
	RoleId uint64 `gorm:"column:role_id;type:bigint(20);not null;comment:'角色ID'"`
	MenuId uint64 `gorm:"column:menu_id;type:bigint(20);not null;comment:'菜单ID'"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
