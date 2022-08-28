package usecase

import (
	"context"
	"strings"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"gorm.io/gorm"
)

// RegisterStudent is the usecase who instantiate a student and save it in the database
func RegisterStudent(
	ctx context.Context,
	repo student.RegisterStudentRepository,
	updateRepo student.UpdateStudentRepository,
	getLoggedUsr auth.GetLoggedUsr,
) student.RegisterStudent {
	userAlreadyIsProducer := func(email string) (baseuser.BaseUser, error) {
		f := filters.NewConditions()
		f.WithCondition("usr_email", email)
		f.WithCondition("usr_scope", baseuser.PRODUCER.String())

		producer, err := getLoggedUsr.Run(ctx, f)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}

		// nil, nil if producer was not found
		return producer, nil
	}

	return func(s student.Student) (student.Student, error) {
		err := s.HashPassword()
		if err != nil {
			return nil, err
		}

		producer, err := userAlreadyIsProducer(s.GetEmail())
		if err != nil {
			return nil, err
		}

		// if producer income as nil it will return false and jump to
		// create a student
		if baseuser.IsProducer(producer) {
			return updateRepo.Run(
				ctx,
				producer.GetID(),
				makesStudent,
			)
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
		return student.ErrEmailAlreadyExists
	}

	return err
}

func makesStudent(s student.Student) (student.Student, error) {
	s.AddScope(baseuser.STUDENT)
	return s, nil
}
