package dao

import (
	"server/internal/dto"
	"server/internal/model"
)

type UserDAO struct{}

func (d *UserDAO) FindList(req dto.UserListReq) []model.UserModel {
	// db.Where(...).Find(&users)
	// return users
	return nil
}

func (d *UserDAO) FindByUsername(username string) *model.UserModel {
	return nil
}
