package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(username string, role string) string {
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		PanicIfError(err)
	}
	return t
}
