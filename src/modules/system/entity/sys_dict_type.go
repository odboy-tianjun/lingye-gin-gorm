package entity

import (
	"lingye-gin/src/base"
)

// 字典类型
type DictType struct {
	base.LingYeDO
	DictName string `gorm:"column:dict_name;type:varchar(255);comment:'字典名称'"`
	DictType string `gorm:"column:dict_type;type:varchar(255);unique_index;comment:'字典类型'"`
	Status   string `gorm:"column:status;type:char(1);default:'0';comment:'状态(0正常 1停用)'"`
	Remark   string `gorm:"column:remark;comment:'备注'"`
}

func (DictType) TableName() string {
	return "sys_dict_type"
}
