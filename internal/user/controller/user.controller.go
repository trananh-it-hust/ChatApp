package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/internal/user/model"
	"main.go/internal/user/service"
	"main.go/pkg/response"
	"main.go/pkg/util"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	user := model.UserRegister{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Invalid request")))
		return
	}
	if user.Username == "" || user.Password == "" || user.Email == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Missing required fields")))
		return
	}
	if !util.IsValidEmail(user.Email) {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Invalid email format")))
		return
	}

	if err := uc.UserService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(err))
		return
	}

	ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "User created successfully", nil))
}

func (uc *UserController) LoginUser(ctx *gin.Context) {
	user := model.UserLogin{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Invalid request")))
		return
	}
	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Missing required fields")))
		return
	}
	if !util.IsValidEmail(user.Email) {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Invalid email format")))
		return
	}
	userModel, err := uc.UserService.LoginUser(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorUnauthorized(err))
		return
	}

	ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Login successful", userModel))
}
