package middleware

import (
	"moviesnow-backend/model/web"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (middleware *AuthMiddleware) VerifyJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userHeader := c.Request().Header["X-Token"]
		if len(userHeader) == 0 {
			// X-Tijeb not passed
			return c.JSON(
				http.StatusUnauthorized,
				web.WebResponse{
					Code:   403,
					Status: "ERROR",
					Data:   "Unauthorized",
				},
			)
		}

		token, err := jwt.Parse(userHeader[0], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				web.WebResponse{
					Code:   403,
					Status: "ERROR",
					Data:   "Unauthorized",
				},
			)
		}

		if !token.Valid {
			return c.JSON(
				http.StatusUnauthorized,
				web.WebResponse{
					Code:   403,
					Status: "ERROR",
					Data:   "Invalid Token",
				},
			)
		}

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
