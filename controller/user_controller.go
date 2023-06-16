package controller

import (
	"moviesnow-backend/helper"
	"moviesnow-backend/model/web"
	"moviesnow-backend/service"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Register(echo.Context) error
	Login(echo.Context) error
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
	err := helper.BindValidate(c, userBody)
	if err != nil {
		helper.PanicIfError(err)
	}

	userResponse, err := userController.UserService.Register(c.Request().Context(), userBody)
	if err != nil {
		if helper.IsMongoDup(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   400,
				Status: "ERROR",
				Data:   "Username has been used",
			})
		}

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

func (userController UserControllerImpl) Login(c echo.Context) error {
	userBody := &web.UserCreateRequest{}
	err := c.Bind(&userBody)
	if err != nil {
		helper.PanicIfError(err)
	}

	userResponse, err := userController.UserService.Login(c.Request().Context(), userBody)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   404,
				Status: "ERROR",
				Data:   "User not found",
			})
		}

		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   401,
				Status: "ERROR",
				Data:   "Wrong password",
			})
		}
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	})
}
