package repository

import (
	"errors"

	"github.com/trananh-it-hust/ChatApp/global"
	"github.com/trananh-it-hust/ChatApp/internal/user/model"
)

type UserRepository interface {
	CreateUser(user model.UserRegister) error
	CheckEmailExist(email string) error
	GetUserByEmail(email string) (model.User, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
func (ur *UserRepositoryImpl) CreateUser(user model.UserRegister) error {
	db := global.MDB
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepositoryImpl) CheckEmailExist(email string) error {
	db := global.MDB
	var count int64
	if err := db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Email already exists")
	}
	return nil
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) (model.User, error) {
	db := global.MDB
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
