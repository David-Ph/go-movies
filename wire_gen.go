// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo"
	"moviesnow-backend/app"
	"moviesnow-backend/controller"
	"moviesnow-backend/repository"
	"moviesnow-backend/service"
)

// Injectors from wire.go:

func InitializeServer() *echo.Echo {
	database := app.NewDB()
	userRepositoryImpl := repository.NewUserRepositoryImpl(database)
	validate := validator.New()
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl, validate)
	userControllerImpl := controller.NewUserControllerImpl(userServiceImpl)
	movieRepositoryImpl := repository.NewMovieRepositoryImpl(database)
	movieServiceImpl := service.NewMovieServiceImpl(movieRepositoryImpl, validate)
	movieControllerImpl := controller.NewMovieControllerImpl(movieServiceImpl)
	reviewRepositoryImpl := repository.NewReviewRepositoryImpl(database)
	reviewServiceImpl := service.NewReviewServiceImpl(reviewRepositoryImpl, validate)
	reviewControllerImpl := controller.NewReviewControllerImpl(reviewServiceImpl)
	router := app.NewRouter(userControllerImpl, movieControllerImpl, reviewControllerImpl)
	echoEcho := app.NewServer(router, database)
	return echoEcho
}

// wire.go:

var UserSet = wire.NewSet(repository.NewUserRepositoryImpl, service.NewUserServiceImpl, controller.NewUserControllerImpl)

var MovieSet = wire.NewSet(repository.NewMovieRepositoryImpl, service.NewMovieServiceImpl, controller.NewMovieControllerImpl)

var ReviewSet = wire.NewSet(repository.NewReviewRepositoryImpl, service.NewReviewServiceImpl, controller.NewReviewControllerImpl)
