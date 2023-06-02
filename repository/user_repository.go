package repository

import (
	"context"
	"fmt"
	"moviesnow-backend/helper"
	"moviesnow-backend/model/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Login(context.Context, *entity.User) *entity.User
	Register(context.Context, *entity.User) *entity.User
	GetUserData(context.Context, *entity.User) *entity.User
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

func (userRepository *UserRepositoryImpl) Login(ctx context.Context, u *entity.User) {
	// Put code here
}

func (userRepository *UserRepositoryImpl) Register(ctx context.Context, u *entity.User) *entity.User {
	user := &entity.User{
		Username: u.Username,
		Password: u.Password,
		Role:     entity.USER,
	}

	res, err := userRepository.DB.Collection("users").InsertOne(ctx, user)
	helper.PanicIfError(err)

	user.Id = fmt.Sprint(res.InsertedID)
	return user
}

func (userRepository *UserRepositoryImpl) GetUserData(ctx context.Context, id string) {
}
