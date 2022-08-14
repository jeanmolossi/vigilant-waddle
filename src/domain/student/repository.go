package student

import "github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"

type GetStudentsRepository interface {
	Run(filters filters.FilterConditions) ([]Student, error)
}
