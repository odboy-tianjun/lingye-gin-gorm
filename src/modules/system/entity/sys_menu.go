package entity

import (
	"lingye-gin/src/base"
)

// 菜单权限
type Menu struct {
	base.LingYeDO
	MenuName       string `gorm:"column:menu_name;type:varchar(255);not null;comment:'菜单名称'"`
	ParentId       uint64 `gorm:"column:parent_id;type:bigint(20);default:0;comment:'父菜单ID'"`
	OrderNum       uint64 `gorm:"column:order_num;type:int(4);default:0;comment:'显示顺序'"`
	Path           string `gorm:"column:path;type:varchar(500);default:'';comment:'路由地址'"`
	Component      string `gorm:"column:component;type:varchar(255);comment:'组件路径'"`
	IsFrame        int    `gorm:"column:is_frame;type:tinyint(1);default:1;comment:'是否为外链(0是 1否)'"`
	IsCache        int    `gorm:"column:is_cache;type:tinyint(1);default:0;comment:'是否缓存(0缓存 1不缓存)'"`
	MenuType       string `gorm:"column:menu_type;type:char(1);default:'';comment:'菜单类型(M目录 C菜单 F按钮)'"`
	Visible        string `gorm:"column:visible;type:char(1);default:'0';comment:'菜单状态(0显示 1隐藏)'"`
	Status         string `gorm:"column:status;type:char(1);default:'0';comment:'菜单状态(0正常 1停用)'"`
	PermissionCode string `gorm:"column:permission_code;type:varchar(100);comment:'权限标识'"`
	Icon           string `gorm:"column:icon;type:varchar(100);default:'#';comment:'菜单图标'"`
}

func (Menu) TableName() string {
	return "sys_menu"
}
