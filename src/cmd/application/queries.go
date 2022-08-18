package application

import (
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type query struct {
	db database.Database
}

func NewQuery(db database.Database) Query {
	return &query{db}
}

func (q *query) GetStudents() student.GetStudents {
	return factory.GetStudents(q.db)
}
