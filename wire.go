//go:build wireinject
// +build wireinject

package main

import (
	"moviesnow-backend/app"
	"moviesnow-backend/controller"
	"moviesnow-backend/repository"
	"moviesnow-backend/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo"
)

var UserSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	service.NewUserServiceImpl,
	controller.NewUserControllerImpl,
)

var MovieSet = wire.NewSet(
	repository.NewMovieRepositoryImpl,
	service.NewMovieServiceImpl,
	controller.NewMovieControllerImpl,
)

var ReviewSet = wire.NewSet(
	repository.NewReviewRepositoryImpl,
	service.NewReviewServiceImpl,
	controller.NewReviewControllerImpl,
)

func InitializeServer() *echo.Echo {
	wire.Build(
		app.NewDB,
		validator.New,
		UserSet,
		MovieSet,
		ReviewSet,
		app.NewRouter,
		app.NewServer,
	)
	return nil
}
