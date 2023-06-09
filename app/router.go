package app

import (
	"moviesnow-backend/controller"

	"github.com/labstack/echo"
)

type Router struct {
	UserController *controller.UserControllerImpl
}

func NewRouter(userController *controller.UserControllerImpl) *Router {
	return &Router{
		UserController: userController,
	}
}

func (r Router) InitializeRoute(e *echo.Echo) {
	userRoute := e.Group("/auth")
	userRoute.POST("/register", r.UserController.Register)
	userRoute.POST("/login", r.UserController.Login)
}
