package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/paginator"
)

type getStudents struct {
	db database.Database
}

func NewGetStudents(db database.Database) student.GetStudentsRepository {
	return &getStudents{db}
}

func (g *getStudents) Run(ctx context.Context, f filters.FilterConditions, p paginator.Paginator) ([]student.Student, error) {
	panic("not implemented repository")
}
