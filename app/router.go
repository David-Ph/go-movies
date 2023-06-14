package app

import (
	"moviesnow-backend/controller"

	"github.com/labstack/echo"
)

type Router struct {
	UserController  *controller.UserControllerImpl
	MovieController *controller.MovieControllerImpl
}

func NewRouter(
	userController *controller.UserControllerImpl,
	movieController *controller.MovieControllerImpl,
) *Router {
	return &Router{
		UserController:  userController,
		MovieController: movieController,
	}
}

func (r Router) InitializeRoute(e *echo.Echo) {
	authRoute := e.Group("/auth")
	authRoute.POST("/register", r.UserController.Register)
	authRoute.POST("/login", r.UserController.Login)

	movieRoute := e.Group("/movie")
	movieRoute.POST("/", r.MovieController.Create)
}
