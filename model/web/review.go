package web

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewCreateRequest struct {
	MovieId primitive.ObjectID `validate:"required" json:"movie_id" form:"movie_id"`
	UserId  primitive.ObjectID `validate:"required" json:"user_id" form:"user_id"`
	Text    string             `validate:"required" json:"text" form:"text"`
	Rating  float64            `validate:"required" json:"rating" form:"rating"`
}

type ReviewResponse struct {
	Id      primitive.ObjectID `json:"id"`
	MovieId primitive.ObjectID `json:"movie_id"`
	UserId  primitive.ObjectID `json:"user_id"`
	Text    string             `json:"text"`
	Rating  float64            `json:"rating"`
}
