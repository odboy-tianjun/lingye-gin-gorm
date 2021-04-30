package base

import (
	"time"
)

type LingYeDO struct {
	// 自增ID
	ID uint64 `gorm:"primary_key"`
	// 创建时间
	CreatedAt time.Time `gorm:"comment:'创建时间'"`
	// 更新时间
	UpdatedAt time.Time `gorm:"comment:'更新时间'"`
	// 删除时间
	DeletedAt *time.Time `gorm:"comment:'删除时间'" sql:"index"`
	// 创建者
	CreatedBy string `gorm:"comment:'创建者'"`
	// 更新者
	UpdatedBy string `gorm:"comment:'更新者'"`
}
