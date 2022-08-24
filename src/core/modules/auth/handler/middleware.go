package handler

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"
	"github.com/labstack/echo/v4"
)

// Middleware is a implementation from echo middleware to check
// if the current user session is valid
func Middleware() echo.MiddlewareFunc {
	db := database.GetConnection()
	usecase := factory.NewValidateAndRefresh(db)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// bypass path who don't need authentication
			if shouldIgnorePath(c.Path()) {
				return next(c)
			}

			var tokenStr string
			token, err := c.Cookie("access_token")
			if err != nil {
				// if the token is not present in cookie try
				// get from Authorization header
				authToken := c.Request().Header.Get("Authorization")
				if authToken == "" {
					return http_error.Handle(c, auth.ErrForbidden)
				}

				tokenStr = authToken
			} else {
				tokenStr = token.Value
			}

			studentID, sessionID, err := Decode(tokenStr)
			if err != nil {
				return http_error.Handle(c, auth.ErrForbidden)
			}

			err = usecase(studentID, sessionID)
			if err != nil {
				c.SetCookie(&http.Cookie{
					Name:    "access_token",
					Path:    "/",
					Expires: time.Unix(0, 0),
				})

				return http_error.Handle(c, auth.ErrForbidden)
			}

			c.Set("studentID", studentID)
			c.Set("sessionID", sessionID)

			log.Println("valid session...")

			return next(c)
		}
	}
}

// shouldIgnorePath receive the current path and check if it will
// bypass by middleware.
//
// If match with anyone path return true.
func shouldIgnorePath(path string) bool {
	middlewareShouldIgnorePaths := []string{
		`^/ping$`,
		`/swagger/(.*)$`,
		`/auth/(login)$`,
		`/student$`,
	}

	for _, p := range middlewareShouldIgnorePaths {
		match, err := regexp.MatchString(p, path)
		if err != nil {
			return false
		}

		if match {
			return true
		}
	}

	return false
}

// Decode receives the token hash and decode that as correctly
// params in order:
//
// 	studentID, sessionID, decodeError
//
// Example:
//
// 	studentID, sessionID, err := Decode(hash)
func Decode(hash string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid session hash")
	}

	return parts[0], parts[1], nil
}
