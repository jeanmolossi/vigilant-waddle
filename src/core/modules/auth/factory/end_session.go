package factory

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func EndSession(
	db database.Database,
) auth.EndSession {
	return usecase.EndSession(
		context.Background(),
		repository.NewDeleteSession(db),
	)
}
