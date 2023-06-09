package controller

import (
	"moviesnow-backend/helper"
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
	userBody := &web.UserCreateRequest{}
	err := c.Bind(&userBody)
	if err != nil {
		helper.PanicIfError(err)
	}

	userResponse, err := userController.UserService.Register(c.Request().Context(), userBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   400,
			Status: "ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	})
}
