package services

import (
	"auth_service/db/repository"
	"auth_service/models"
	"auth_service/utils"
	"fmt"
	env "auth_service/config/env"
	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	GetByID() (*models.User, error)
	CreateUser() error
	LoginUser() (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (user *UserServiceImpl) GetByID() (*models.User, error) {
	fmt.Println("Fetching user by ID")
	userResult, err := user.userRepository.GetByID()
	if err != nil {
		fmt.Println("Error fetching user by ID", err)
		return nil, err
	}
	return userResult, nil
}

func (user *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user")
	encrpytedPassword, err := utils.HashPassword("test")
	if err != nil {
		fmt.Println("Error hashing password", err)
		return err
	}
	user.userRepository.CreateUser("user", "username@gamil.com", encrpytedPassword)
	return nil
}

func (user *UserServiceImpl) LoginUser() (string,error) {
	fmt.Println("Login process for a user")
	userResult, err := user.userRepository.GetUserByEmail("username@gamil.com")
	if(err != nil){
		fmt.Println("Error fetching user by email", err)
		return "", err
	}
	if userResult == nil {
		fmt.Println("User not found")
		return "", err
	}
	doesMatch := utils.CheckPasswordHash("test", userResult.Password)
	if(!doesMatch){
		fmt.Println("Password does not match")
		return "", err
	}

	payload := jwt.MapClaims{
		"email" : userResult.Email,
		"id" : userResult.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "secret")))
	if err != nil {
		fmt.Println("Error signing token", err)
		return "", err
	}
	fmt.Println("Printing the token : ", tokenString)
	return tokenString, nil
}

