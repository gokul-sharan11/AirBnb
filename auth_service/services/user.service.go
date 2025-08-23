package services

import (
	"auth_service/db/repository"
	"auth_service/models"
	"auth_service/utils"
	"fmt"
	env "auth_service/config/env"
	"github.com/golang-jwt/jwt/v4"
	"auth_service/dto"
)

type UserService interface {
	GetByID(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (user *UserServiceImpl) GetByID(id string) (*models.User, error) {
	fmt.Println("Fetching user by ID")
	userResult, err := user.userRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error fetching user by ID", err)
		return nil, err
	}
	return userResult, nil
}

func (user *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user")
	encrpytedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password", err)
		return nil, err
	}

	userResult, err := user.userRepository.CreateUser(payload.Username, payload.Email, encrpytedPassword)
	if err != nil {
		fmt.Println("Error creating user", err)
		return nil, err
	}
	return userResult, nil
}

func (user *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string,error) {
	fmt.Println("Login process for a user")
	userResult, err := user.userRepository.GetUserByEmail(payload.Email)
	if(err != nil){
		fmt.Println("Error fetching user by email", err)
		return "", err
	}
	if userResult == nil {
		fmt.Println("User not found")
		return "", err
	}
	doesMatch := utils.CheckPasswordHash(payload.Password, userResult.Password)
	fmt.Println("Does Match : ", doesMatch)
	if(!doesMatch){
		fmt.Println("Password does not match")
		return "", err
	}

	jwtPayload := jwt.MapClaims{
		"email" : userResult.Email,
		"id" : userResult.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "secret")))
	if err != nil {
		fmt.Println("Error signing token", err)
		return "", err
	}
	fmt.Println("Printing the token : ", tokenString)
	return tokenString, nil
}

