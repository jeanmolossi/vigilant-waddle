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
		return repo.Run(ctx, s)
	}
}
