package controller

import (
	"moviesnow-backend/model/web"
	"moviesnow-backend/service"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	Register(c echo.Context) error
}

type UserControllerImpl struct {
	UserService *service.UserServiceImpl
}

func NewUserControllerImpl(userService *service.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController UserControllerImpl) Register(c echo.Context) error {
	userBody := &web.UserCreateRequest{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	userResponse := userController.UserService.Register(c.Request().Context(), userBody)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}