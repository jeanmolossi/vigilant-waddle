package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/adapter"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"

	"github.com/labstack/echo/v4"
)

// RegitersStudent godoc
//
// @Summary Register a student
// @Description Register a student
// @ID register-student
// @Produce json
// @Param student body adapter.RegisterStudent true "Student"
// @Success 201 {object} student.Student
// @Failure 400 {object} http_error.HTTPBadRequestError
// @Failure 409 {object} http_error.HTTPError
// @Failure 500 {object} http_error.HTTPError
// @Router /student [post]
func RegisterStudent() echo.HandlerFunc {
	db := database.GetConnection()
	usecase := factory.RegisterStudent(db)

	return func(c echo.Context) error {
		studentInput := new(adapter.RegisterStudent)

		if err := c.Bind(studentInput); err != nil {
			return http_error.Handle(c, err)
		}

		if err := c.Validate(studentInput); err != nil {
			return http_error.Handle(c, err)
		}

		s := student.NewStudent(
			student.WithEmail(studentInput.Email),
			student.WithPassword(studentInput.Password),
		)

		err := usecase(s)

		if err != nil {
			return http_error.Handle(c, err)
		}

		return c.JSON(http.StatusCreated, s)
	}
}
