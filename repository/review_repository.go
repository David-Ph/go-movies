package repository

import (
	"context"
	"moviesnow-backend/model/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepository interface {
	Create(context.Context, *entity.Review, primitive.ObjectID, primitive.ObjectID) (*entity.Review, error)
	Update(context.Context, *entity.Review) (entity.Review, error)
	Delete(context.Context) ([]*entity.Review, error)
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

	// result := reviewRepository.DB.Collection("movies").FindOneAndUpdate(ctx, bson.M{})

	res, err := reviewRepository.DB.Collection("reviews").InsertOne(ctx, review)
	if err != nil {
		return nil, err
	}

	review.Id = res.InsertedID.(primitive.ObjectID)
	return review, nil
}
