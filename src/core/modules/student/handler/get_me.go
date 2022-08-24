package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"
	"github.com/labstack/echo/v4"
)

// GetMe godoc
//
// @Summary Get current student
// @Description Get current student
// @ID get-me
// @Tags student
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
		input := new(student.GetMeOptions)

		if err := c.Bind(input); err != nil {
			return http_error.Handle(c, err)
		}

		input.StudentID = c.Get("studentID").(string)

		if err := c.Validate(input); err != nil {
			return http_error.Handle(c, err)
		}

		student, err := usecase(*input)
		if err != nil {
			return http_error.Handle(c, err)
		}

		if student != nil {
			return c.JSON(http.StatusOK, NewHttpNewStudent(student))
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
