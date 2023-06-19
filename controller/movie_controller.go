package controller

import (
	"moviesnow-backend/helper"
	"moviesnow-backend/model/web"
	"moviesnow-backend/service"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieController interface {
	Create(echo.Context) error
	FindAll(echo.Context) error
	FindById(echo.Context) error
	GetCategories(echo.Context) error
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
	params := &web.MovieFilterQuery{}
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

func (movieController MovieControllerImpl) FindById(c echo.Context) error {
	movieId := c.Param("movie_id")

	if movieId == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   400,
			Status: "ERROR",
			Data:   "Movie ID is empty",
		})
	}

	movieResponse, err := movieController.MovieService.FindById(c.Request().Context(), movieId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   404,
				Status: "ERROR",
				Data:   "Movie not found",
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
		Data:   movieResponse,
	})
}

func (movieController MovieControllerImpl) GetCategories(c echo.Context) error {

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   helper.Categories,
	})
}
