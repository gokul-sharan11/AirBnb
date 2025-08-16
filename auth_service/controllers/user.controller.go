package controllers

import (
	"net/http"
	"auth_service/services"
	"fmt"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	u.userService.GetByID()
	w.Write([]byte("User registered"))
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user using the request object")
	u.userService.CreateUser()
	w.Write([]byte("User reated"))
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	u.userService.LoginUser()
	w.Write([]byte("User logged in"))
}