package setup

import (
	"context"
	"fmt"
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
	name, err := db.Collection("users").Indexes().CreateOne(context.TODO(), indexModel)
	fmt.Println(name)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}
}
