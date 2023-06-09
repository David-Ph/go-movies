package service

import (
	"context"
	"moviesnow-backend/helper"
	"moviesnow-backend/model/entity"
	"moviesnow-backend/model/web"
	"moviesnow-backend/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(context.Context, *entity.User) *entity.User
	Register(context.Context, *entity.User) *entity.User
	GetUserData(context.Context, *entity.User) *entity.User
}

type UserServiceImpl struct {
	UserRepository *repository.UserRepositoryImpl
	Validate       *validator.Validate
}

// Constructor for UserServiceImpl
func NewUserServiceImpl(userRepository *repository.UserRepositoryImpl, validate *validator.Validate) *UserServiceImpl {
	o := &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
	return o
}

// func (userService *UserServiceImpl) Login(ctx context.Context, u *entity.User) *entity.User {
// 	// Put code here
// }

func (userService *UserServiceImpl) Register(ctx context.Context, request *web.UserCreateRequest) (*web.UserResponse, error) {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := &entity.User{
		Username: request.Username,
		Password: string(hashedPassword),
	}

	user, err = userService.UserRepository.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	token := helper.SignToken(user.Username, string(user.Role))

	return &web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}, nil
}

// func (userService *UserServiceImpl) GetUserData(ctx context.Context, u *entity.User) *entity.User {
// 	// Put code here
// }
