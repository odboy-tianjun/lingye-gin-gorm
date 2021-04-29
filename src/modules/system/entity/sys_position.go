package entity

import (
	"lingye-gin/src/base"
)

// 岗位
type Position struct {
	base.LingYeDO
	PostCode string `gorm:"column:post_code;type:varchar(255);not null;comment:'岗位编码'"`
	PostName string `gorm:"column:post_name;type:varchar(255);not null;comment:'岗位名称'"`
	PostSort int    `gorm:"column:post_sort;type:int(4);not null;comment:'显示顺序'"`
	Status   string `gorm:"column:status;type:char(1);not null;comment:'状态(0正常 1停用)'"`
	Remark   string `gorm:"column:remark;type:varchar(255);comment:'备注'"`
}

func (Position) TableName() string {
	return "sys_position"
}
