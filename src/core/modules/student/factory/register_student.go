package factory

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func RegisterStudent(
	db database.Database,
) student.RegisterStudent {
	return usecase.RegisterStudent(
		context.Background(),
		repository.NewRegisterStudent(db),
	)
}
