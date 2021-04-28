package entity

import (
	"lingye-gin/src/base"
)

type User struct {
	base.LingYeDO
	UserName    string `gorm:"column:user_name;type:varchar(255);not null;index:idx_username"`
	NickName    string `gorm:"column:nick_name;type:varchar(255)"`
	Email       string `gorm:"column:email;type:varchar(255)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255)"`
	Sex         string `gorm:"column:sex;type:varchar(255)"`
	Avatar      string `gorm:"column:avatar;type:varchar(255)"`
	Password    string `gorm:"column:password;type:varchar(255);not null"`
	Status      string `gorm:"column:status;type:varchar(255);not null"`
	Remark      string `gorm:"column:remark;type:text(0)"`
}

// 设置User的表名为`sys_user`
func (User) TableName() string {
	return "sys_user"
}
