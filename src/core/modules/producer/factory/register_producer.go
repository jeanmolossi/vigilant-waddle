package factory

import (
	"context"

	authRepository "github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/producer/repository"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/producer/usecase"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func RegisterProducer(
	db database.Database,
) producer.RegisterProducer {
	return usecase.RegisterProducer(
		context.Background(),
		repository.NewRegisterProducer(db),
		repository.NewUpdateProducer(db),
		authRepository.NewGetLoggedUser(db),
	)
}
