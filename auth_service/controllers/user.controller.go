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
	userId := r.URL.Query().Get("id")
	if userId == "" {
		userId = r.Context().Value("userID").(string)
	}

	fmt.Println("Fetching user by ID")

	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	user, err := u.userService.GetByID(userId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", nil)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully")
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user using the request object")

    payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

	user, err := u.userService.CreateUser(&payload)

	if err != nil {
		fmt.Println("Error creating user", err)
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User created successfully", user)
	fmt.Println("User created successfully")
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	jwtToken, err := u.userService.LoginUser(&payload)
	if err != nil {
		fmt.Println("Error logging in user", err)
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}