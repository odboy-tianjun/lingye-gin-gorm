package entity

import (
	"lingye-gin/src/base"
)

// 部门
type Dept struct {
	base.LingYeDO
	ParentId uint64 `gorm:"column:parent_id;type:bigint(20);default:0;comment:'父部门ID'"`
	DeptName string `gorm:"column:dept_name;type:varchar(255);default:'';comment:'部门名称'"`
	OrderNum int    `gorm:"column:order_num;type:int(4);default:0;comment:'显示顺序'"`
	LeaderId uint64 `gorm:"column:leader;type:varchar(20);comment:'负责人ID'"`
	Status   string `gorm:"column:status;type:char(1);default:'0';comment:'部门状态(0正常 1停用)'"`
}

func (Dept) TableName() string {
	return "sys_dept"
}
