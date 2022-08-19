// Package usecase will implements all student usecases
package usecase

import (
	"context"
	"fmt"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/paginator"
)

// GetMe is a usecase to get a student
//
// It will return a student or an error
func GetMe(
	ctx context.Context,
	repo student.GetStudentsRepository,
) student.GetMe {
	return func(gmo student.GetMeOptions) (student.Student, error) {
		f := filters.NewConditions()
		f.AddFields(gmo.Fields)
		f.WithCondition("usr_id", gmo.StudentID)

		p := paginator.New()

		students, err := repo.Run(ctx, f, p)
		if err != nil {
			return nil, err
		}

		if len(students) == 0 {
			return nil, fmt.Errorf("student not found")
		}

		return students[0], nil
	}
}
