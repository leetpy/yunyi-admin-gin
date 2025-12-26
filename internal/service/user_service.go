package service

import (
	"context"
	"errors"
	"fmt"
	"server/internal/dao"
	"server/internal/dto"
	"server/internal/infra"
	"server/internal/model"
	"server/internal/pkg/apperror"
	"server/internal/pkg/token"
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

func (s *UserService) Login(req dto.UserLoginReq) (*dto.UserLoginResp, error) {
	userInfo, err := s.userDao.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if !s.CheckPassword(req.Password, userInfo.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成 token
	jwtCfg := infra.AppConfig.Jwt
	accessToken, refreshTokne, err := token.GenerateTokenPair(*userInfo, jwtCfg.AccessSignedKey, jwtCfg.RefreshSignedKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil, apperror.New(apperror.BuildTokenErr, "生成token失败")
	}

	// 将 token 存入 Redis 白名单
	ctx := context.Background()
	accessTokenKey := fmt.Sprintf("whitelist:access_token:%s", accessToken)
	refreshTokenKey := fmt.Sprintf("whitelist:refresh_token:%s", refreshTokne)

	// accessToken 过期时间 24 小时
	if infra.RDS != nil {
		infra.RDS.Set(ctx, accessTokenKey, userInfo.ID, time.Hour*24)
		// refreshToken 过期时间 7 天
		infra.RDS.Set(ctx, refreshTokenKey, userInfo.ID, time.Hour*24*7)
	}

	// 更新最后登录时间
	userInfo.LastLoginAt.Scan(time.Now())
	go s.userDao.Update(userInfo)

	return &dto.UserLoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshTokne,
		User: struct {
			Username string `json:"username"`
			UserId   uint   `json:"user_id"`
		}{
			Username: userInfo.Username,
			UserId:   userInfo.ID,
		},
	}, nil
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
