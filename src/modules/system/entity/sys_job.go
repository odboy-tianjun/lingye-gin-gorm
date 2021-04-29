package entity

import (
	"lingye-gin/src/base"
)

// 定时任务
type Job struct {
	base.LingYeDO
	JobName        string `gorm:"column:job_name;type:varchar(255);not null;comment:'任务名称'"`
	InvokeTarget   string `gorm:"column:invoke_target;type:varchar(255);not null;comment:'调用目标字符串'"`
	CronExpression string `gorm:"column:cron_expression;type:varchar(255);not null;comment:'cron执行表达式'"`
	MisfirePolicy  string `gorm:"column:misfire_policy;type:varchar(255);not null;default:'3';comment:'计划执行错误策略(1立即执行 2执行一次 3放弃执行)'"`
	Concurrent     string `gorm:"column:concurrent;type:varchar(255);not null;default:'1';comment:'是否并发执行(0允许 1禁止)'"`
	Status         string `gorm:"column:status;type:char(1);not null;default:'0';comment:'状态(0正常 1暂停)'"`
}

func (Job) TableName() string {
	return "sys_job"
}
