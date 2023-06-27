package service

import (
	"context"
	"errors"
	"moviesnow-backend/model/entity"
	"moviesnow-backend/model/web"
	"moviesnow-backend/repository"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewService interface {
	Create(context.Context, *web.ReviewCreateRequest) (*entity.Review, error)
	Delete(context.Context, primitive.ObjectID) (*entity.Review, error)
	FindByMovieId(context.Context, primitive.ObjectID) ([]*entity.Review, error)
	FindByUserId(context.Context, primitive.ObjectID) ([]*entity.Review, error)
}

type ReviewServiceImpl struct {
	ReviewRepository *repository.ReviewRepositoryImpl
	Validate         *validator.Validate
}

func NewReviewServiceImpl(reviewRepository *repository.ReviewRepositoryImpl, validate *validator.Validate) *ReviewServiceImpl {
	o := &ReviewServiceImpl{
		ReviewRepository: reviewRepository,
		Validate:         validate,
	}
	return o
}

func (reviewService *ReviewServiceImpl) Create(ctx context.Context, request *web.ReviewCreateRequest) (*web.ReviewResponse, error) {
	err := reviewService.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	review := &entity.Review{
		MovieId: request.MovieId,
		UserId:  request.UserId,
		Text:    request.Text,
		Rating:  request.Rating,
	}

	review, err = reviewService.ReviewRepository.Create(ctx, review)
	if err != nil {
		return nil, err
	}

	return &web.ReviewResponse{
		Id:      review.Id,
		MovieId: review.MovieId,
		UserId:  review.UserId,
		Text:    review.Text,
		Rating:  review.Rating,
	}, nil
}

func (reviewService *ReviewServiceImpl) Delete(ctx context.Context, id string) (*web.ReviewResponse, error) {
	reviewId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid review id")
	}

	_, err = reviewService.ReviewRepository.Delete(ctx, reviewId)
	if err != nil {
		return nil, err
	}

	return &web.ReviewResponse{}, nil
}

func (reviewService *ReviewServiceImpl) FindByMovieId(ctx context.Context, id string) ([]entity.Review, error) {
	movieId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid movie id")
	}

	reviews, err := reviewService.ReviewRepository.FindByMovieId(ctx, movieId)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (reviewService *ReviewServiceImpl) FindByUserId(ctx context.Context, id string) ([]entity.Review, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	reviews, err := reviewService.ReviewRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
