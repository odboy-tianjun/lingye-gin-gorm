package entity

import (
	"lingye-gin/src/base"
)

// 角色
type Role struct {
	base.LingYeDO
	RoleName  string `gorm:"column:role_name;type:varchar(30);not null;comment:'角色名称'"`
	RoleKey   string `gorm:"column:role_key;type:varchar(100);not null;comment:'角色权限字符串'"`
	RoleSort  int    `gorm:"column:role_sort;type:int(4);not null;default:0;comment:'显示顺序'"`
	DataScope string `gorm:"column:data_scope;type:char(1);not null;default:'1';comment:'数据范围(1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限)'"`
	Status    string `gorm:"column:status;type:char(1);not null;default:'0';comment:'角色状态(0正常 1停用)'"`
}

func (Role) TableName() string {
	return "sys_role"
}
