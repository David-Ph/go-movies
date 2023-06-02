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

func InitializeServer() *echo.Echo {
	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewUserRepositoryImpl,
		service.NewUserServiceImpl,
		controller.NewUserControllerImpl,
		app.NewRouter,
		app.NewServer,
	)
	return nil
}
