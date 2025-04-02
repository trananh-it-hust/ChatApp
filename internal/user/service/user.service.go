package service

import (
	"errors"

	"main.go/internal/user/model"
	"main.go/internal/user/repository"
	"main.go/pkg/util"
)

type UserService interface {
	CreateUser(user model.UserRegister) error
	LoginUser(user model.UserLogin) (model.UserLoginResponse, error)
}
type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService() UserService {
	return &UserServiceImpl{
		UserRepository: repository.NewUserRepository(),
	}
}
func (s *UserServiceImpl) CreateUser(user model.UserRegister) error {
	if err := s.UserRepository.CheckEmailExist(user.Email); err != nil {
		return err
	}
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := s.UserRepository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) LoginUser(user model.UserLogin) (model.UserLoginResponse, error) {
	userModel, err := s.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		return model.UserLoginResponse{}, err
	}
	if !util.CheckPasswordHash(user.Password, userModel.Password) {
		return model.UserLoginResponse{}, errors.New("Invalid password")
	}

	token, err := util.GenerateToken(userModel.ID, userModel.Email)
	if err != nil {
		return model.UserLoginResponse{}, err
	}
	return model.UserLoginResponse{
		Token:    token,
		Username: userModel.Username,
		Email:    userModel.Email,
	}, nil
}
