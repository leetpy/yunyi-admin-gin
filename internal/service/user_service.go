package service

import (
	"errors"
	"server/internal/dao"
	"server/internal/dto"
	"server/internal/infra"
	"server/internal/model"
	"server/internal/pkg/apperror"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDao *dao.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDAO(infra.DB),
	}
}

func (s *UserService) Login(req dto.UserLoginReq) (*model.UserModel, error) {
	userInfo, err := s.userDao.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if !s.CheckPassword(req.Password, userInfo.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 更新最后登录时间
	userInfo.LastLoginAt = time.Now()
	go s.userDao.Update(userInfo)

	return userInfo, nil
}

func (s *UserService) GetUserByUsername(username string) (*model.UserModel, error) {
	return s.userDao.FindByUsername(username)
}

func (s *UserService) CreateUser(req dto.CreateUserReq) (*model.UserModel, error) {
	var err error
	req.Password, err = s.HashPassword(req.Password)
	if err != nil {
		return nil, apperror.New(apperror.PasswordWrong, "无效的密码")
	}

	userInfo, err := s.userDao.CreateUser(req)
	if err != nil {
		return nil, apperror.New(apperror.UserExists, "用户已存在")
	}

	return userInfo, nil
}

// 密码hash
func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 密码比对
func (s *UserService) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
