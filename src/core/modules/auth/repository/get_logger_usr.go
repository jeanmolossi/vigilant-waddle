package repository

import (
	"context"
	"errors"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

type getLoggedUsr struct {
	db database.Database
}

func NewGetLoggedUser(db database.Database) auth.GetLoggedUsr {
	return &getLoggedUsr{db}
}

func (g *getLoggedUsr) Run(ctx context.Context, f filters.FilterConditions) (baseuser.BaseUser, error) {
	result := g.db.DB().Table("users")

	if !f.HasConditions() {
		return nil, errors.New("you should provide user id")
	}

	if fields, ok := f.WithFields("users"); ok {
		result = result.Select(fields)
	}

	statement, values := f.Conditions()
	result = result.Where(statement, values...)

	user := new(UserModel)
	result = result.First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	isStudent := user.Type == baseuser.STUDENT.String()

	if isStudent {
		return student.NewStudent(
			student.WithID(user.ID),
			student.WithEmail(user.Email),
			student.WithScope(user.Type),
		), nil
	}

	return producer.NewProducer(
		producer.WithID(user.ID),
		producer.WithEmail(user.Email),
		producer.WithScope(user.Type),
	), nil
}
