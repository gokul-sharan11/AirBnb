package router

import (
	"auth_service/controllers"
	"auth_service/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (u *UserRouter) Register(r chi.Router) {
	r.Get("/profile", u.userController.GetUserById)
	r.With(middlewares.UserCreateRequestValidator).Post("/register", u.userController.CreateUser)
	r.With(middlewares.UserLoginRequestValidator).Post("/login", u.userController.LoginUser)
}