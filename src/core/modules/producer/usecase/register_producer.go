package usecase

import (
	"context"
	"strings"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
)

// RegisterProducer is the usecase who instantiate a producer and save it in the database
func RegisterProducer(
	ctx context.Context,
	repo producer.RegisterProducerRepository,
) producer.RegisterProducer {
	return func(s producer.Producer) (producer.Producer, error) {
		err := s.HashPassword()
		if err != nil {
			return nil, err
		}

		st, err := repo.Run(ctx, s)
		if err != nil {
			return nil, catchDuplicateErr(err)
		}

		return st, nil
	}
}

// catchDuplicateErr will catch the duplicate error and return a custom error
func catchDuplicateErr(err error) error {
	if strings.Contains(err.Error(), "duplicate key") {
		return producer.ErrEmailAlreadyExists
	}

	return err
}
