package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type registerStudent struct {
	db database.Database
}

func NewRegisterStudent(db database.Database) student.RegisterStudentRepository {
	return &registerStudent{db}
}

func (r *registerStudent) Run(ctx context.Context, s student.Student) error {
	student := &StudentModel{
		ID:       s.GetID(),
		Email:    s.GetEmail(),
		Password: s.GetPassword(),
	}

	return r.db.DB().Create(student).Error
}
