package controllers

import (
	"auth_service/dto"
	"auth_service/services"
	"auth_service/utils"
	"fmt"
	"net/http"
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

	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong", jsonErr)
		return
	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}

	jwtToken, err := u.userService.LoginUser(&payload)
	if err != nil {
		fmt.Println("Error logging in user", err)
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}