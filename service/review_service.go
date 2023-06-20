package service

import (
	"context"
	"moviesnow-backend/model/entity"
	"moviesnow-backend/model/web"
	"moviesnow-backend/repository"

	"github.com/go-playground/validator/v10"
)

type ReviewService interface {
	Create(context.Context, *web.ReviewCreateRequest) (*entity.Review, error)
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
