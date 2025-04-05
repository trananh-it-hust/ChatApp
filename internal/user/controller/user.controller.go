package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trananh-it-hust/ChatApp/internal/user/model"
	"github.com/trananh-it-hust/ChatApp/internal/user/service"
	"github.com/trananh-it-hust/ChatApp/pkg/response"
	"github.com/trananh-it-hust/ChatApp/pkg/util"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserRegister true "User register"
// @Success 200 {object} response.Response{data=null}
// @Failure 400 {object} response.Response
// @Router /auth/register [post]
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
