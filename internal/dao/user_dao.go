package dao

import (
	"server/internal/dto"
	"server/internal/model"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (d *UserDAO) FindList(req dto.UserListReq) []model.UserModel {
	// db.Where(...).Find(&users)
	// return users
	return nil
}

func (d *UserDAO) CreateUser(req dto.CreateUserReq) (*model.UserModel, error) {
	user := model.UserModel{
		Username: req.Username,
		Password: req.Password,
		State:    model.USER_STATE_ENABLED,
	}
	err := d.db.Create(&user).Error
	return &user, err
}

func (d *UserDAO) FindByUsername(username string) (*model.UserModel, error) {
	var user model.UserModel
	err := d.db.First(&user).Where("username = ?", username).Error
	return &user, err
}

func (d *UserDAO) Update(user *model.UserModel) error {
	return d.db.Save(user).Error
}
