package dao

import (
	"server/internal/dto"
	"server/internal/model"
)

type UserDAO struct{}

func (d *UserDAO) FindList(req dto.UserListReq) []model.User {
	// db.Where(...).Find(&users)
	// return users
	return nil
}
