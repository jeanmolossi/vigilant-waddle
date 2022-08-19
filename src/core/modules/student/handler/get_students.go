package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/labstack/echo/v4"
)

// GetStudents godoc
//
// @Summary Get all students
// @Description Get all students
// @ID get-students
// @Tags student
// @Produce json
// @Success 200 {object} []student.Student
// @Security access_token
// @Router /students [get]
func GetStudents() echo.HandlerFunc {
	db := database.GetConnection()

	usecase := factory.GetStudents(db)

	return func(c echo.Context) error {

		students, err := usecase()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		if len(students) > 0 {
			return c.JSON(http.StatusOK, students)
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
