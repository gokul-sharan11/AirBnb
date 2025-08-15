package router

import (
	"github.com/go-chi/chi/v5"
	"auth_service/controllers"
)

type UserRouter struct {
	userController controllers.UserController
}

func NewUserRouter(_userController controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (u *UserRouter) Register(r chi.Router) {
	r.Post("/register", u.userController.RegisterUser)
}