package usecase

import (
	"context"
	"strings"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
)

// RegisterStudent is the usecase who instantiate a student and save it in the database
func RegisterStudent(
	ctx context.Context,
	repo student.RegisterStudentRepository,
) student.RegisterStudent {
	return func(s student.Student) (student.Student, error) {
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
		return student.ErrEmailAlreadyExists
	}

	return err
}
