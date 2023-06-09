package app

import (
	"moviesnow-backend/setup"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewServer(r *Router, db *mongo.Database) *echo.Echo {
	e := echo.New()

	setup.UserCollectionIndex(db)
	r.InitializeRoute(e)

	return e
}
