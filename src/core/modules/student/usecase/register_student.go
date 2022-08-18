package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
)

func RegisterStudent(
	ctx context.Context,
	repo student.RegisterStudentRepository,
) student.RegisterStudent {
	return func(s student.Student) error {
		err := s.HashPassword()
		if err != nil {
			return err
		}

		return repo.Run(ctx, s)
	}
}
