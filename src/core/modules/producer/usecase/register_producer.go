package usecase

import (
	"context"
	"strings"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"gorm.io/gorm"
)

// RegisterProducer is the usecase who instantiate a producer and save it in the database
func RegisterProducer(
	ctx context.Context,
	repo producer.RegisterProducerRepository,
	updateRepo producer.UpdateProducerRepository,
	getLoggerUsr auth.GetLoggedUsr,
) producer.RegisterProducer {
	userAlreadyIsStudent := func(email string) (baseuser.BaseUser, error) {
		f := filters.NewConditions()
		f.WithCondition("usr_email", email)
		f.WithCondition("usr_scope", baseuser.STUDENT.String())

		student, err := getLoggerUsr.Run(ctx, f)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}

		// nil, nil if student was not found
		return student, nil
	}

	return func(p producer.Producer) (producer.Producer, error) {
		err := p.HashPassword()
		if err != nil {
			return nil, err
		}

		student, err := userAlreadyIsStudent(p.GetEmail())
		if err != nil {
			return nil, err
		}

		// if student income as nil it will return false and jump to
		// create a producer
		if baseuser.IsStudent(student) {
			return updateRepo.Run(
				ctx,
				student.GetID(),
				makesProducer,
			)
		}

		st, err := repo.Run(ctx, p)
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

func makesProducer(p producer.Producer) (producer.Producer, error) {
	p.AddScope(baseuser.PRODUCER)
	return p, nil
}
