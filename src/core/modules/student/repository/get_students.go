package repository

import (
	"context"
	"log"

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
	var dbStudents []StudentModel

	log.Println("DB", g.db.DB())

	result := g.db.DB().Find(&dbStudents)

	if result.Error != nil {
		return nil, result.Error
	}

	students := make([]student.Student, len(dbStudents))
	for i, s := range dbStudents {
		students[i] = student.NewStudent(
			student.WithID(s.ID),
			student.WithEmail(s.Email),
		)
	}

	return students, nil
}
