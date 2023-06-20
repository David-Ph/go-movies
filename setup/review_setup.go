package setup

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReviewCollectionIndex(db *mongo.Database) {
	//unique index here
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "movie_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Collection("reviews").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}
}
