package service

import (
	"context"
	"moviesnow-backend/model/entity"
	"moviesnow-backend/model/web"
	"moviesnow-backend/repository"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type MovieService interface {
	Create(context.Context, *web.MovieCreateRequest) (*entity.Movie, error)
	FindAll(context.Context) ([]*entity.Movie, error)
	FindById(context.Context, string) (entity.Movie, error)
}

type MovieServiceImpl struct {
	MovieRepository *repository.MovieRepositoryImpl
	Validate        *validator.Validate
}

func NewMovieServiceImpl(movieRepository *repository.MovieRepositoryImpl, validate *validator.Validate) *MovieServiceImpl {
	o := &MovieServiceImpl{
		MovieRepository: movieRepository,
		Validate:        validate,
	}
	return o
}

func (movieService *MovieServiceImpl) Create(ctx context.Context, request *web.MovieCreateRequest) (*web.MovieResponse, error) {
	err := movieService.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	posterUrl, err := url.ParseRequestURI(request.Poster)
	if err != nil {
		return nil, err
	}

	movie := &entity.Movie{
		Title:      request.Title,
		Poster:     posterUrl.String(),
		Rating:     0.0,
		Info:       request.Info,
		Categories: request.Categories,
	}

	movie, err = movieService.MovieRepository.Create(ctx, movie)
	if err != nil {
		return nil, err
	}

	return &web.MovieResponse{
		Id:         movie.Id,
		Title:      movie.Title,
		Poster:     movie.Poster,
		Rating:     movie.Rating,
		Info:       movie.Info,
		Categories: movie.Categories,
	}, nil
}

func (movieService *MovieServiceImpl) FindAll(ctx context.Context, query *web.MovieFilterQuery) ([]entity.Movie, error) {
	err := movieService.Validate.Struct(query)
	if err != nil {
		return nil, err
	}

	movies, err := movieService.MovieRepository.FindAll(ctx, query)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (movieService *MovieServiceImpl) FindById(ctx context.Context, movieId string) (*web.MovieResponse, error) {
	m, err := movieService.MovieRepository.FindById(ctx, movieId)
	if err != nil {
		return nil, err
	}

	return &web.MovieResponse{
		Title:      m.Title,
		Poster:     m.Poster,
		Rating:     m.Rating,
		Info:       m.Info,
		Categories: m.Categories,
		Id:         m.Id,
	}, nil
}
