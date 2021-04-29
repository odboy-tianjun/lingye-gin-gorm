package entity

import (
	"lingye-gin/src/base"
)

// 操作日志
type OperateLog struct {
	base.LingYeDO
	Title           string `gorm:"column:title;type:varchar(255);comment:'模块标题'"`
	BusinessType    string `gorm:"column:business_type;type:char(1);comment:'业务类型(0其它 1新增 2修改 3删除)'"`
	Method          string `gorm:"column:method;type:varchar(100);comment:'方法名称'"`
	RequestMethod   string `gorm:"column:request_method;type:varchar(10);comment:'请求方式'"`
	DeptName        string `gorm:"column:dept_name;type:varchar(50);comment:'部门名称'"`
	OperateUrl      string `gorm:"column:operate_url;type:varchar(255);comment:'请求url'"`
	OperateIp       string `gorm:"column:operate_ip;type:varchar(255);comment:'主机地址'"`
	OperateLocation string `gorm:"column:operate_location;type:varchar(255);comment:'操作地点'"`
	OperateParam    string `gorm:"column:operate_param;type:varchar(255);comment:'请求参数'"`
	JsonResult      string `gorm:"column:json_result;type:varchar(255);comment:'返回参数'"`
	Status          string `gorm:"column:status;type:char(1);comment:'操作状态(0正常 1异常)'"`
	ErrorMsg        string `gorm:"column:error_msg;type:text(0);comment:'错误消息'"`
}

func (OperateLog) TableName() string {
	return "sys_operate_log"
}
