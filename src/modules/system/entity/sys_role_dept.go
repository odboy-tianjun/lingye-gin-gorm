package entity

// 角色和部门关联
type RoleDept struct {
	RoleId uint64 `gorm:"column:role_id;type:bigint(20);not null;comment:'角色ID'"`
	DeptId uint64 `gorm:"column:dept_id;type:bigint(20);not null;comment:'部门ID'"`
}

func (RoleDept) TableName() string {
	return "sys_role_dept"
}
