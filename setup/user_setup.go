package setup

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserCollectionIndex(db *mongo.Database) {
	//unique index here
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Collection("users").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}
}
