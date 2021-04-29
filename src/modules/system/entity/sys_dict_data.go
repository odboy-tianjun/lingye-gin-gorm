package entity

import (
	"lingye-gin/src/base"
)

// 字典数据
type DictData struct {
	base.LingYeDO
	DictSort  int    `gorm:"column:dict_sort;type:int(4);default:0;comment:'字典排序'"`
	DictLabel string `gorm:"column:dict_label;type:varchar(255);comment:'字典标签'"`
	DictValue string `gorm:"column:dict_value;type:varchar(255);comment:'字典键值'"`
	DictType  string `gorm:"column:dict_type;type:varchar(255);comment:'字典类型'"`
	IsDefault string `gorm:"column:is_default;type:varchar(255);default:'N';comment:'是否默认(Y是 N否)'"`
	Status    string `gorm:"column:status;type:char(1);default:'0';comment:'状态(0正常 1停用)'"`
	Remark    string `gorm:"column:remark;type:varchar(255);comment:'字典标签'"`
}

func (DictData) TableName() string {
	return "sys_dict_data"
}
