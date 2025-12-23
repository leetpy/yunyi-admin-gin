package model

type UserModel struct {
	Username string `gorm:"uniqueIndex;not null; comment:用户名"`
	Password string `gorm:"not null; comment:密码"`
}

func (UserModel) TableName() string {
	return "user"
}
