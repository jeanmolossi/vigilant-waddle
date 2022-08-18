package application

import "github.com/jeanmolossi/vigilant-waddle/src/domain/student"

type Option func(*app) error

type Application interface {
	Query() Query
	Command() Command
}

type Query interface {
	GetStudents() student.GetStudents
}

type Command interface{}
