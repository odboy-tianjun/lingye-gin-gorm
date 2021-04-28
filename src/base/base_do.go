package base

import (
	"time"
)

type BaseDO struct {
	// 自增ID
	ID uint `gorm:"column:id;type:bigint(20);not null;primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
	// 删除时间
	DeletedAt *time.Time `sql:"index"`
	// 扩展字段
	// 创建者
	CreatedBy string
	// 删除者
	UpdatedBy string
}
