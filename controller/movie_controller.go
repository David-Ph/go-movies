package controller

import (
	"moviesnow-backend/helper"
	"moviesnow-backend/model/web"
	"moviesnow-backend/service"
	"net/http"

	"github.com/labstack/echo"
)

type MovieController interface {
	Create(echo.Context) error
	FindAll(echo.Context) error
	FindById(echo.Context) error
}

type MovieControllerImpl struct {
	MovieService *service.MovieServiceImpl
}

func NewMovieControllerImpl(movieService *service.MovieServiceImpl) *MovieControllerImpl {
	return &MovieControllerImpl{
		MovieService: movieService,
	}
}

func (movieController MovieControllerImpl) Create(c echo.Context) error {
	movieBody := &web.MovieCreateRequest{}
	err := helper.BindValidate(c, movieBody)
	if err != nil {
		helper.PanicIfError(err)
	}

	movieResponse, err := movieController.MovieService.Create(c.Request().Context(), movieBody)
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
		Data:   movieResponse,
	})
}

func (movieController MovieControllerImpl) FindAll(c echo.Context) error {
	params := &web.MovieFilterParam{}
	err := helper.BindValidate(c, params)
	if err != nil {
		helper.PanicIfError(err)
	}

	movieResponse, err := movieController.MovieService.FindAll(c.Request().Context(), params)
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
		Data:   movieResponse,
	})
}
