package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Info struct {
	ReleaseDate  string `bson:"release_date" json:"release_date"`
	Director     string `bson:"director" json:"director"`
	FeaturedSong string `bson:"featured_song" json:"featured_song"`
}

type Movie struct {
	Id     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Poster string             `bson:"poster"`
	Title  string             `bson:"title"`
	Rating float64            `bson:"rating"`
	Info   Info               `bson:"info"`
}
