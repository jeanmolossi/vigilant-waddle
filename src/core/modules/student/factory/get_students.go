package factory

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func GetStudents(
	db database.Database,
) student.GetStudents {
	return usecase.GetStudents(
		context.Background(),
		repository.NewGetStudents(db),
	)
}
