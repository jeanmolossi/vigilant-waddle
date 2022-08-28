package student

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/paginator"
)

// GetStudentsRepository is the repository implementation for Get Students.
type GetStudentsRepository interface {
	// Run executes the Get Students query.
	//
	// Example:
	//
	//      students, err := repo.Run(context.Backround(), filters.New(), paginator.NewConditions())
	//
	Run(ctx context.Context, f filters.FilterConditions, p paginator.Paginator) ([]Student, error)
}

// RegisterStudentRepository is the repository implementation for Register Student.
type RegisterStudentRepository interface {
	// Run executes the Register Student query.
	//
	// Example:
	//
	//      newStudent, err := repo.Run(context.Backround(), student)
	//
	Run(ctx context.Context, s Student) (Student, error)
}

// StudentUpdater is a interface who defines the UpdateStudentRepository method
type StudentUpdater func(Student) (Student, error)

// UpdateStudentRepository is repository implementation for Update Student.
type UpdateStudentRepository interface {
	// Run executes the update in student
	Run(ctx context.Context, usrID string, u StudentUpdater) (Student, error)
}
