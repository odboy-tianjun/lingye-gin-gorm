package entity

import (
	"lingye-gin/src/base"
)

// 用户信息
type User struct {
	base.LingYeDO
	UserName    string `gorm:"column:user_name;type:varchar(255);index:idx_username;comment:'用户真名'"`
	UserCode    string `gorm:"column:user_code;type:varchar(255);unique_index;comment:'用户编码'"`
	NickName    string `gorm:"column:nick_name;type:varchar(255);comment:'用户昵称'"`
	Email       string `gorm:"column:email;type:varchar(255);unique_index;comment:'邮箱'"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255);unique_index;comment:'手机号'"`
	Sex         string `gorm:"column:sex;type:varchar(255);comment:'性别(0男 1女 2未知)'"`
	Avatar      string `gorm:"column:avatar;type:varchar(255);comment:'头像路径'"`
	Password    string `gorm:"column:password;type:varchar(255);comment:'加密密码'"`
	Status      string `gorm:"column:status;type:char(1);default:'0';comment:'帐号状态(0正常 1停用)'"`
	Remark      string `gorm:"column:remark;type:text(0);comment:'备注'"`
}

// 设置User的表名为`sys_user`
func (User) TableName() string {
	return "sys_user"
}
