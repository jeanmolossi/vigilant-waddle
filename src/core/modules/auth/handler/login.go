package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/adapter"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"
	"github.com/labstack/echo/v4"
)

// Login handles the login request.
// It validates the credentials and creates a new session.
// It returns a JSON with the access token.
//
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body adapter.AuthCredentials true "Login credentials"
// @Success 200 {object} HttpAccessToken
// @Failure 400 {object} http_error.HTTPError
// @Failure 401 {object} http_error.HTTPError
// @Failure 403 {object} http_error.HTTPError
// @Failure 500 {object} http_error.HTTPError
// @Router /auth/login [post]
func Login() echo.HandlerFunc {
	db := database.GetConnection()
	usecase := factory.StartSession(db)

	return func(c echo.Context) error {
		credentialsInput := new(adapter.AuthCredentials)

		if err := c.Bind(credentialsInput); err != nil {
			return http_error.Handle(c, err)
		}

		if err := c.Validate(credentialsInput); err != nil {
			return http_error.Handle(c, err)
		}

		s, err := usecase(
			credentialsInput.Email,
			credentialsInput.Password,
		)

		if err != nil {
			return http_error.Handle(c, err)
		}

		if s.GetID() == "" {
			return http_error.Handle(c, student.ErrInvalidCredentials)
		}

		c.SetCookie(&http.Cookie{
			Name:    "access_token",
			Value:   s.HashToken(),
			Path:    "/",
			Expires: s.GetRefreshToken().Expiration(),
		})

		return c.JSON(http.StatusOK, HttpAccessToken{
			AccessToken: s.HashToken(),
		})
	}
}

// Http responses

// HttpAccessToken is a access token representation for http response
type HttpAccessToken struct {
	AccessToken string `json:"access_token" example:"OGE4MTlhMTctYTMxZS00OTE0LWE4ZjAtMzQ1Njg5ZThiMzg1OjJmZjhiOGIzLTU0OWItNGRjMi04Mjc4LWVhMDdlNjQxMGY1ZA=="`
}
