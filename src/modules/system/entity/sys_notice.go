package entity

import (
	"lingye-gin/src/base"
)

// 通知公告
type Notice struct {
	base.LingYeDO
	NoticeTitle   string `gorm:"column:notice_title;type:varchar(50);not null;comment:'公告标题'"`
	NoticeType    string `gorm:"column:notice_type;type:char(1);not null;comment:'公告类型(1通知 2公告)'"`
	NoticeContent string `gorm:"column:notice_content;type:text(0);not null;comment:'公告内容'"`
	Status        string `gorm:"column:status;type:char(1);not null;default:'0';comment:'公告状态(0正常 1关闭)'"`
	Remark        string `gorm:"column:remark"`
}

func (Notice) TableName() string {
	return "sys_notice"
}
