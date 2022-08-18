package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/paginator"
)

func GetStudents(
	ctx context.Context,
	repo student.GetStudentsRepository,
) student.GetStudents {
	return func() ([]student.Student, error) {
		f := filters.NewConditions()
		p := paginator.New()
		return repo.Run(ctx, f, p)
	}
}
