package service

import (
	"server/internal/dao"
	"server/internal/dto"
	"server/internal/model"
)

type UserService struct{}

func (s *UserService) GetList(req dto.UserListReq) []model.User {
	return dao.UserDAO.FindList(req)
}
