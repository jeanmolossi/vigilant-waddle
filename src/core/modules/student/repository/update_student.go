package repository

import (
	"context"
	"errors"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type updateStudent struct {
	db database.Database
}

// NewUpdateStudent will return a student.UpdateStudentRepository implementation
func NewUpdateStudent(db database.Database) student.UpdateStudentRepository {
	return &updateStudent{db}
}

// Run executes the update in student
func (u *updateStudent) Run(ctx context.Context, usrID string, updater student.StudentUpdater) (student.Student, error) {
	// when has not user should fail.
	// that method only works if has usrID
	if usrID == "" {
		return nil, errors.New("fail on get usr with ID empty")
	}

	model := &StudentModel{
		ID: usrID,
	}

	result := u.db.DB().First(model)
	if result.Error != nil {
		return nil, result.Error
	}

	// update the user that was
	s, err := updater(model.AsDomain())
	if err != nil {
		return nil, err
	}

	updatedModel := &StudentModel{
		ID:        s.GetID(),
		Email:     s.GetEmail(),
		Password:  s.GetPassword(),
		Type:      s.GetScope().String(),
		CreatedAt: model.CreatedAt,
	}

	result = u.db.DB().Save(updatedModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return updatedModel.AsDomain(), nil
}
