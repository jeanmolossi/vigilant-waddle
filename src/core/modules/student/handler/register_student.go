package handler

import (
	"net/http"

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
// @Success 201 {object} student.Student
// @Router /student [post]
func RegisterStudent() echo.HandlerFunc {
	db := database.GetConnection()

	usecase := factory.RegisterStudent(db)

	return func(c echo.Context) error {
		s := student.NewStudent(
			student.WithEmail(c.QueryParam("email")),
			student.WithPassword(c.QueryParam("password")),
		)

		err := usecase(s)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, s)
	}
}
