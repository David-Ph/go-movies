package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	MovieId primitive.ObjectID `bson:"movie_id"`
	UserId  primitive.ObjectID `bson:"user_id"`
	Rating  float64            `bson:"rating"`
	Text    string             `bson:"text"`
}
