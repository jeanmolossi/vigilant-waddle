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
// @Tags student
// @Produce json
// @Param student body adapter.RegisterStudent true "Student"
// @Success 201 {object} HttpNewStudent
// @Failure 400 {object} http_error.HTTPBadRequestError "Bad request"
// @Failure 409 {object} http_error.HTTPError "User with that email already exists"
// @Failure 500 {object} http_error.HTTPError "An error occurred"
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

		student, err := usecase(s)

		if err != nil {
			return http_error.Handle(c, err)
		}

		return c.JSON(http.StatusCreated, NewHttpNewStudent(student))
	}
}

// Http responses

// HttpStudent is a student representation for http response
type HttpStudent struct {
	ID    string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Email string `json:"email" example:"john@doe.com" format:"email"`
}

// HttpNewStudent is a student representation for http response
type HttpNewStudent struct {
	Data HttpStudent `json:"data"`
}

// NewHttpNewStudent creates a new HttpNewStudent
func NewHttpNewStudent(s student.Student) *HttpNewStudent {
	return &HttpNewStudent{
		Data: HttpStudent{s.GetID(), s.GetEmail()},
	}
}
