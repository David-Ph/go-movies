package app

import (
	"moviesnow-backend/controller"
	"moviesnow-backend/middleware"

	"github.com/labstack/echo"
)

type Router struct {
	UserController   *controller.UserControllerImpl
	MovieController  *controller.MovieControllerImpl
	ReviewController *controller.ReviewControllerImpl
	AuthMiddleware   *middleware.AuthMiddleware
}

func NewRouter(
	userController *controller.UserControllerImpl,
	movieController *controller.MovieControllerImpl,
	reviewController *controller.ReviewControllerImpl,
	authMiddleware *middleware.AuthMiddleware,
) *Router {
	return &Router{
		UserController:   userController,
		MovieController:  movieController,
		ReviewController: reviewController,
		AuthMiddleware:   authMiddleware,
	}
}

func (r Router) InitializeRoute(e *echo.Echo) {
	authRoute := e.Group("/auth")
	authRoute.POST("/register", r.UserController.Register)
	authRoute.POST("/login", r.UserController.Login)

	movieRoute := e.Group("/movie")
	movieRoute.POST("", r.MovieController.Create)
	movieRoute.GET("", r.MovieController.FindAll)
	movieRoute.GET("/:movie_id", r.MovieController.FindById)
	movieRoute.GET("/categories", r.MovieController.GetCategories)

	reviewRoute := e.Group("/review")
	reviewRoute.POST("", r.ReviewController.Create, r.AuthMiddleware.VerifyJWT)
	reviewRoute.DELETE("/:id", r.ReviewController.Delete)
	reviewRoute.GET("/user/:user_id", r.ReviewController.FindByUserId)
	reviewRoute.GET("/movie/:movie_id", r.ReviewController.FindByMovieId)

}
