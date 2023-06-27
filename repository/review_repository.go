package repository

import (
	"context"
	"moviesnow-backend/helper"
	"moviesnow-backend/model/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepository interface {
	Create(context.Context, *entity.Review, primitive.ObjectID, primitive.ObjectID) (*entity.Review, error)
	Delete(context.Context) ([]*entity.Review, error)
	FindByMovieId(context.Context, primitive.ObjectID) ([]*entity.Review, error)
	FindByUserId(context.Context, primitive.ObjectID) ([]*entity.Review, error)
}

type ReviewRepositoryImpl struct {
	DB *mongo.Database
}

func NewReviewRepositoryImpl(db *mongo.Database) *ReviewRepositoryImpl {
	o := &ReviewRepositoryImpl{
		DB: db,
	}
	return o
}

func (reviewRepository *ReviewRepositoryImpl) Create(ctx context.Context, r *entity.Review) (*entity.Review, error) {
	review := &entity.Review{
		Id:      primitive.NewObjectID(),
		UserId:  r.UserId,
		MovieId: r.MovieId,
		Text:    r.Text,
		Rating:  r.Rating,
	}

	// ? create review
	res, err := reviewRepository.DB.Collection("reviews").InsertOne(ctx, review)
	if err != nil {
		return nil, err
	}

	review.Id = res.InsertedID.(primitive.ObjectID)

	// ? Find average of movie rating
	matchStage := bson.D{
		{Key: "$match", Value: bson.M{"movie_id": r.MovieId}},
	}
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$movie_id"},
			{Key: "averageRating", Value: bson.D{{Key: "$avg", Value: "$rating"}}},
		}},
	}

	reviewsCursor, err := reviewRepository.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = reviewsCursor.All(ctx, &results); err != nil {
		helper.PanicIfError(err)
	}

	// ? Update the reviewed movie's update
	update := bson.M{
		"$set": bson.M{"rating": results[0]["averageRating"]},
	}
	_, _ = reviewRepository.DB.Collection("movies").UpdateByID(ctx, r.MovieId, update)

	return review, nil
}

func (reviewRepository *ReviewRepositoryImpl) Delete(ctx context.Context, id primitive.ObjectID) (*entity.Review, error) {
	reviews := &entity.Review{}
	filter := bson.M{
		"_id": id,
	}

	_, err := reviewRepository.DB.Collection("reviews").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (reviewRepository *ReviewRepositoryImpl) FindByMovieId(ctx context.Context, movieId primitive.ObjectID) ([]entity.Review, error) {
	reviews := []entity.Review{}
	filter := bson.M{
		"movie_id": movieId,
	}

	cursor, err := reviewRepository.DB.Collection("reviews").Find(ctx, filter)
	if err != nil {
		helper.PanicIfError(err)
	}
	if err = cursor.All(ctx, &reviews); err != nil {
		helper.PanicIfError(err)
	}

	return reviews, nil
}

func (reviewRepository *ReviewRepositoryImpl) FindByUserId(ctx context.Context, userId primitive.ObjectID) ([]entity.Review, error) {
	reviews := []entity.Review{}
	filter := bson.M{
		"user_id": userId,
	}

	cursor, err := reviewRepository.DB.Collection("reviews").Find(ctx, filter)
	if err != nil {
		helper.PanicIfError(err)
	}
	if err = cursor.All(ctx, &reviews); err != nil {
		helper.PanicIfError(err)
	}

	return reviews, nil
}
