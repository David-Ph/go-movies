//go:build wireinject
// +build wireinject

package main

import (
	"moviesnow-backend/app"

	"github.com/google/wire"
	"github.com/labstack/echo"
)

func InitializeServer() *echo.Echo {
	wire.Build(
		app.NewDB,
		app.NewRouter,
		app.NewServer,
	)
	return nil
}
