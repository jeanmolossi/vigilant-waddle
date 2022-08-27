package factory

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func NewValidateAndRefresh(db database.Database) auth.ValidateAndRefresh {
	return usecase.ValidateAndRefresh(
		context.Background(),
		repository.NewGetSingleSession(db),
		repository.NewUpdateSession(db),
		repository.NewGetLoggedUser(db),
	)
}
