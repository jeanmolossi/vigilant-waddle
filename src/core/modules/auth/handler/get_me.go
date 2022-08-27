package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"
	"github.com/labstack/echo/v4"
)

// GetMe godoc
//
// @Summary Get current student
// @Description Get current student
// @ID get-me
// @Tags auth
// @Produce json
// @Param fields query []string false "fields to return from the student"
// @Success 200 {object} HttpNewStudent
// @Failure 403 {object} http_error.HTTPError "Forbidden"
// @Failure 404 {object} http_error.HTTPError "User not found"
// @Failure 500 {object} http_error.HTTPError "An error occurred"
// @Security access_token
// @Router /me [get]
func GetMe() echo.HandlerFunc {
	db := database.GetConnection()

	usecase := factory.GetMe(db)

	return func(c echo.Context) error {
		input := new(auth.GetMeOptions)

		if c.Get("studentID").(string) != "" {
			input.UserID = c.Get("studentID").(string)
		}

		if c.Get("producerID").(string) != "" {
			input.UserID = c.Get("producerID").(string)
		}

		if err := c.Bind(input); err != nil {
			return http_error.Handle(c, err)
		}

		if err := c.Validate(input); err != nil {
			return http_error.Handle(c, err)
		}

		user, err := usecase(*input)
		if err != nil {
			return http_error.Handle(c, err)
		}

		if user != nil {
			return c.JSON(http.StatusOK, NewHttpUser(user))
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}

// Responses

// HttpUser is a student representation for http response
type HttpUser struct {
	ID    string `json:"id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Email string `json:"email,omitempty" example:"john@doe.com" format:"email"`
	Scope string `json:"scopes,omitempty" example:"student"`
}

// HttpNewUser is a student representation for http response
type HttpNewUser struct {
	Data HttpUser `json:"data"`
}

// NewHttpUser creates a new HttpNewUser
func NewHttpUser(s baseuser.BaseUser) *HttpNewUser {
	return &HttpNewUser{
		Data: HttpUser{
			s.GetID(),
			s.GetEmail(),
			string(s.GetScope()),
		},
	}
}
