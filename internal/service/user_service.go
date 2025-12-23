package service

import (
	"server/internal/dto"
	"server/internal/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(req dto.UserListReq) []model.UserModel {
	return nil
	// return dao.UserDAO.FindList(req)
}

// func (s *UserService) GetList(req dto.UserListReq) []model.User {
// 	return dao.UserDAO.FindList(req)
// }
