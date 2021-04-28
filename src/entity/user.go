package entity

import (
	"lingye-gin/src/base"
)

type User struct {
	base.BaseDO
	UserName    string `gorm:"column:user_name;type:varchar(255);not null;index:idx_username" json:"username" form:"username"`
	NickName    string `gorm:"column:nick_name;type:varchar(255)" json:"nickname" form:"nickname"`
	Email       string `gorm:"column:email;type:varchar(255)" json:"email" form:"email"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255)" json:"phoneNumber" form:"phoneNumber"`
	Sex         string `gorm:"column:sex;type:varchar(255)" json:"sex" form:"sex"`
	Avatar      string `gorm:"column:avatar;type:varchar(255)" json:"avatar" form:"avatar"`
	Password    string `gorm:"column:password;type:varchar(255);not null" json:"password" form:"password"`
	Status      string `gorm:"column:status;type:varchar(255);not null" json:"status" form:"status"`
	Remark      string `gorm:"column:remark;type:text(0)" json:"remark" form:"remark"`
}

// 设置User的表名为`sys_user`
func (User) TableName() string {
	return "sys_user"
}
