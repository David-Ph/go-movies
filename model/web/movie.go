package web

import (
	"moviesnow-backend/model/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieCreateRequest struct {
	Title      string      `validate:"required,max=200,min=1" json:"title" form:"title"`
	Poster     string      `json:"poster" form:"poster"`
	Info       entity.Info `json:"info" form:"info"`
	Categories []string    `json:"categories" form:"categories"`
}

type MovieFilterQuery struct {
	Categories string `query:"categories" validate:"omitempty"`
	Page       int64  `query:"page" validate:"omitempty,number"`
	Limit      int64  `query:"limit" validate:"omitempty,number"`
	Skip       int64  `query:"skip" validate:"omitempty,number"`
}

type MovieResponse struct {
	Id         primitive.ObjectID `json:"id"`
	Title      string             `json:"title"`
	Poster     string             `json:"poster"`
	Rating     float64            `json:"rating"`
	Info       entity.Info        `json:"info"`
	Categories []string           `json:"categories"`
}
