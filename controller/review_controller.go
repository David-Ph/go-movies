package controller

import (
	"moviesnow-backend/helper"
	"moviesnow-backend/model/web"
	"moviesnow-backend/service"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewController interface {
	Create(echo.Context) error
	Delete(echo.Context) error
	FindByMovieId(echo.Context) error
	FindByUserId(echo.Context) error
}

type ReviewControllerImpl struct {
	ReviewService *service.ReviewServiceImpl
}

func NewReviewControllerImpl(reviewService *service.ReviewServiceImpl) *ReviewControllerImpl {
	return &ReviewControllerImpl{
		ReviewService: reviewService,
	}
}

func (reviewController ReviewControllerImpl) Create(c echo.Context) error {
	reviewBody := &web.ReviewCreateRequest{}
	err := helper.BindValidate(c, reviewBody)
	if err != nil {
		helper.PanicIfError(err)
	}

	reviewResponse, err := reviewController.ReviewService.Create(c.Request().Context(), reviewBody)
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
		Data:   reviewResponse,
	})
}

func (reviewController ReviewControllerImpl) Delete(c echo.Context) error {
	reviewId := c.Param("review_id")

	if reviewId == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   400,
			Status: "ERROR",
			Data:   "Review ID is empty",
		})
	}

	reviewResponse, err := reviewController.ReviewService.Delete(c.Request().Context(), reviewId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   404,
				Status: "ERROR",
				Data:   "Review not found",
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
		Data:   reviewResponse,
	})
}

func (reviewController ReviewControllerImpl) FindByMovieId(c echo.Context) error {
	movieId := c.Param("movie_id")

	if movieId == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   400,
			Status: "ERROR",
			Data:   "Movie ID is empty",
		})
	}

	reviewResponse, err := reviewController.ReviewService.FindByMovieId(c.Request().Context(), movieId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   404,
				Status: "ERROR",
				Data:   "Review not found",
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
		Data:   reviewResponse,
	})
}

func (reviewController ReviewControllerImpl) FindByUserId(c echo.Context) error {
	userId := c.Param("user_id")

	if userId == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   400,
			Status: "ERROR",
			Data:   "Movie ID is empty",
		})
	}

	reviewResponse, err := reviewController.ReviewService.FindByUserId(c.Request().Context(), userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   404,
				Status: "ERROR",
				Data:   "Review not found",
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
		Data:   reviewResponse,
	})
}
