package repository

import (
	"context"
	"fmt"
	"moviesnow-backend/helper"
	"moviesnow-backend/model/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	Create(context.Context, *entity.Movie) (*entity.Movie, error)
	FindAll(context.Context) ([]*entity.Movie, error)
	FindById(context.Context, string) (entity.Movie, error)
}

type MovieRepositoryImpl struct {
	DB *mongo.Database
}

func NewMovieRepositoryImpl(db *mongo.Database) *MovieRepositoryImpl {
	o := &MovieRepositoryImpl{
		DB: db,
	}
	return o
}

func (movieRepository *MovieRepositoryImpl) Create(ctx context.Context, m *entity.Movie) (*entity.Movie, error) {
	movie := &entity.Movie{
		Id:         primitive.NewObjectID(),
		Title:      m.Title,
		Poster:     m.Poster,
		Rating:     m.Rating,
		Info:       m.Info,
		Categories: m.Categories,
	}

	res, err := movieRepository.DB.Collection("movies").InsertOne(ctx, movie)
	if err != nil {
		return nil, err
	}

	movie.Id = res.InsertedID.(primitive.ObjectID)
	fmt.Println(movie)
	return movie, nil
}

func (movieRepository *MovieRepositoryImpl) FindAll(ctx context.Context, filter interface{}) ([]entity.Movie, error) {
	result := []entity.Movie{}

	cursor, err := movieRepository.DB.Collection("movies").Find(ctx, filter)
	if err != nil {
		helper.PanicIfError(err)
	}
	if err = cursor.All(ctx, &result); err != nil {
		helper.PanicIfError(err)
	}

	return result, nil
}

// // Implements MovieRepository
// // TODO: Comment Here
// func (movieRepositoryImpl *MovieRepositoryImpl) FindById(context.Context context.Context, string string) (entity.Movie, error) {
// 	// Put code here
// }
