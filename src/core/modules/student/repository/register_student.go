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

func (r *registerStudent) Run(ctx context.Context, s student.Student) (student.Student, error) {
	_student := &StudentModel{
		ID:       s.GetID(),
		Email:    s.GetEmail(),
		Password: s.GetPassword(),
	}

	result := r.db.DB().Create(_student)
	if result.Error != nil {
		return nil, result.Error
	}

	newStudent := student.NewStudent(
		student.WithID(_student.ID),
		student.WithEmail(_student.Email),
	)

	return newStudent, nil
}
