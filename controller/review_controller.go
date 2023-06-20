package controller

import (
	"moviesnow-backend/helper"
	"moviesnow-backend/model/web"
	"moviesnow-backend/service"
	"net/http"

	"github.com/labstack/echo"
)

type ReviewController interface {
	Create(echo.Context) error
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
