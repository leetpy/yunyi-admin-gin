package model

import "time"

const (
	USER_STATE_ENABLED = iota
	USER_STATE_DISABLED
)

type UserModel struct {
	BaseModel
	Username    string    `gorm:"column:username"`
	Password    string    `gorm:"column:password"`
	State       int       `gorm:"column:state"`
	LastLoginAt time.Time `gorm:"column:last_login_at"`
}

func (UserModel) TableName() string {
	return "users"
}
