package entity

// 用户与岗位关联
type UserPosition struct {
	UserId     uint64 `gorm:"column:user_id;type:bigint(20);not null;comment:'用户ID'"`
	PositionId uint64 `gorm:"column:post_id;type:bigint(20);not null;comment:'岗位ID'"`
}

func (UserPosition) TableName() string {
	return "sys_user_position"
}
