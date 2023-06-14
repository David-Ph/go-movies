package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole string

const (
	ADMIN UserRole = "ADMIN"
	USER  UserRole = "USER"
)

type User struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	Role       UserRole           `bson:"role"`
	Watchlists []Movie            `bson:"watchlists"`
}
