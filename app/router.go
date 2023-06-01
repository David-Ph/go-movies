package app

import (
	"context"
	"log"
	"moviesnow-backend/model/entity"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r Router) InitializeRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		db, err := NewDB()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		csr, err := db.Collection("user").Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err.Error())
		}
		defer csr.Close(ctx)

		result := make([]entity.User, 0)
		for csr.Next(ctx) {
			var row entity.User
			err := csr.Decode(&row)
			if err != nil {
				log.Fatal(err.Error())
			}

			result = append(result, row)
		}

		return c.JSON(http.StatusOK, result)
	})

	e.POST("/", func(c echo.Context) error {
		db, err := NewDB()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err = db.Collection("user").InsertOne(ctx, entity.User{
			Username: "MaoMao",
			Password: "Password",
			Role:     entity.ADMIN,
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		return c.String(http.StatusOK, "Insert Success")
	})
}
