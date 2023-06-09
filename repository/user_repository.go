package repository

import (
	"context"
	"moviesnow-backend/model/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByUsername(context.Context, *entity.User) (*entity.User, error)
	Register(context.Context, *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
	DB *mongo.Database
}

func NewUserRepositoryImpl(db *mongo.Database) *UserRepositoryImpl {
	o := &UserRepositoryImpl{
		DB: db,
	}
	return o
}

func (userRepository *UserRepositoryImpl) FindByUsername(ctx context.Context, u *entity.User) (*entity.User, error) {
	result := &entity.User{}
	res := userRepository.DB.Collection("users").FindOne(ctx, bson.D{
		{Key: "username", Value: u.Username},
	})

	err := res.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (userRepository *UserRepositoryImpl) Register(ctx context.Context, u *entity.User) (*entity.User, error) {
	user := &entity.User{
		Id:       primitive.NewObjectID(),
		Username: u.Username,
		Password: u.Password,
		Role:     entity.USER,
	}

	res, err := userRepository.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}
