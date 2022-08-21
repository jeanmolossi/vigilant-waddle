package handler

import (
	"net/http"
	"time"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"
	"github.com/labstack/echo/v4"
)

// Logout handles the logout request.
// It deletes the session.
// It returns a JSON with the access token.
//
// @Summary Logout
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Success 202 {object} HttpAcceptedLogout
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security access_token
// @Router /auth/logout [post]
func Logout() echo.HandlerFunc {
	db := database.GetConnection()
	usecase := factory.EndSession(db)

	return func(c echo.Context) error {
		sessionID := c.Get("sessionID").(string)

		err := usecase(sessionID)
		if err != nil {
			return http_error.Handle(c, err)
		}

		c.SetCookie(&http.Cookie{
			Name:    "access_token",
			Value:   "null",
			Expires: time.Unix(0, 0),
		})

		return c.JSON(http.StatusAccepted, HttpAcceptedLogout{
			Message: "logged out",
		})
	}
}

// Http Responses

// HttpAcceptedLogout is the response for the logout request.
type HttpAcceptedLogout struct {
	Message string `json:"message" example:"logged out"`
}
