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
type RegisterStudent func(Student) (Student, error)
