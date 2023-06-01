package app

import "github.com/labstack/echo"

func NewServer(r *Router) *echo.Echo {
	e := echo.New()

	r.InitializeRoute(e)

	return e
}
