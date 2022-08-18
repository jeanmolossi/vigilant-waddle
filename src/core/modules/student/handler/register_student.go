package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/adapter"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
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
// @Router /student [post]
func RegisterStudent() echo.HandlerFunc {
	db := database.GetConnection()
	usecase := factory.RegisterStudent(db)

	return func(c echo.Context) error {
		studentInput := new(adapter.RegisterStudent)

		if err := c.Bind(studentInput); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := c.Validate(studentInput); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		s := student.NewStudent(
			student.WithEmail(studentInput.Email),
			student.WithPassword(studentInput.Password),
		)

		err := usecase(s)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, s)
	}
}
