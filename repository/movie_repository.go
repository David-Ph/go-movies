package repository

import (
	"context"
	"fmt"
	"moviesnow-backend/helper"
	"moviesnow-backend/model/entity"
	"moviesnow-backend/model/web"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func (movieRepository *MovieRepositoryImpl) FindAll(ctx context.Context, query *web.MovieFilterQuery) ([]entity.Movie, error) {
	result := []entity.Movie{}
	skip := int64(query.Page*query.Limit - query.Limit)
	options := options.FindOptions{Limit: &query.Limit, Skip: &skip}
	filter := bson.M{}

	if query.Categories != "" {
		filter = bson.M{
			"categories": cases.Title(language.English, cases.Compact).String(query.Categories),
		}
	}

	cursor, err := movieRepository.DB.Collection("movies").Find(ctx, filter, &options)
	if err != nil {
		helper.PanicIfError(err)
	}
	if err = cursor.All(ctx, &result); err != nil {
		helper.PanicIfError(err)
	}

	return result, nil
}

func (movieRepository *MovieRepositoryImpl) FindById(ctx context.Context, movieId primitive.ObjectID) (*entity.Movie, error) {
	result := &entity.Movie{}
	filter := bson.M{
		"_id": movieId,
	}
	res := movieRepository.DB.Collection("movies").FindOne(ctx, filter)

	err := res.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
