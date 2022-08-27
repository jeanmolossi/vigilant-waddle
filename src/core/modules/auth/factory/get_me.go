package factory

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func GetMe(
	db database.Database,
) auth.GetMe {
	return usecase.GetMe(
		context.Background(),
		repository.NewGetLoggedUser(db),
	)
}
