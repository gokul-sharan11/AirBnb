package services

import (
	"auth_service/db/repository"
	"fmt"
)

type UserService interface {
	GetByID() error
	CreateUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (user *UserServiceImpl) GetByID() error {
	fmt.Println("Fetching user by ID")
	user.userRepository.GetByID()
	return nil
}

func (user *UserServiceImpl) CreateUser() error {
	fmt.Println("Fetching user by ID")
	user.userRepository.CreateUser()
	return nil
}

