package service

import (
	"context"
	"moviesnow-backend/helper"
	"moviesnow-backend/model/entity"
	"moviesnow-backend/model/web"
	"moviesnow-backend/repository"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func (userService *UserServiceImpl) Register(ctx context.Context, request *web.UserCreateRequest) *web.UserResponse {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := &entity.User{
		Username: request.Username,
		Password: string(hashedPassword),
	}

	user = userService.UserRepository.Register(ctx, user)

	jwtKey := os.Getenv("JWT_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		helper.PanicIfError(err)
	}

	return &web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    t,
	}
}

// func (userService *UserServiceImpl) GetUserData(ctx context.Context, u *entity.User) *entity.User {
// 	// Put code here
// }
