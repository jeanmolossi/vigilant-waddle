package factory

import (
	"context"

	srepo "github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/repository"

	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
)

func StartSession(
	db database.Database,
) auth.StartSession {
	return usecase.StartSession(
		context.Background(),
		repository.NewCreateSession(db),
		srepo.NewGetStudents(db),
	)
}
