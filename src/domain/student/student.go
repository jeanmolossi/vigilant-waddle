// Package student is a domain package wich contains the student features and behavior.
package student

import baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"

// Student is has the baseUser interface implemented
type Student interface {
	baseuser.BaseUser
}

// Usecases is the student usecases

// GetStudents is the usecase to get all students
// Not implemented yet
type GetStudents func() ([]Student, error)

// RegisterStudent is the usecase to register a new student
type RegisterStudent func(Student) error

// GetMe is the usecase to get the current student
// It will use the session token
type GetMe func(GetMeOptions) (Student, error)

// GetMeOptions is the options for the GetMe usecase
type GetMeOptions struct {
	StudentID string   `json:"student_id" example:"1" format:"uuid" validate:"required,uuid"`
	Fields    []string `query:"fields" example:"name,email"`
}

// GetErrorMap implements ModuleErrorMap to GetMeOptions
func (g *GetMeOptions) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"studentid": {
			"required": ErrMissingStudentID,
			"uuid":     ErrStudentIDInvalid},
	}
}
