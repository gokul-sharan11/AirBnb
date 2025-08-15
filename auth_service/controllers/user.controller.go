package controllers

import (
	"net/http"
	"auth_service/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	
	u.userService.CreateUser()
	w.Write([]byte("User registered"))
}