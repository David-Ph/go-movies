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
	err := c.Bind(&movieBody)
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
