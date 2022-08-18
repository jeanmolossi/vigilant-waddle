// Package student is a domain package wich contains the student features and behavior.
package student

import baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"

// Student is has the baseUser interface implemented
type Student interface {
	baseuser.BaseUser
}

type GetStudents func() ([]Student, error)

type RegisterStudent func(Student) error
