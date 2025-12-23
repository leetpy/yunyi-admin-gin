package model

type UserModel struct{}

func (UserModel) TableName() string {
	return "user"
}
