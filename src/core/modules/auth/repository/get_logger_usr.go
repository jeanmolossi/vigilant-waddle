package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type getLoggedUsr struct {
	db database.Database
}

func NewGetLoggedUser(db database.Database) auth.GetLoggedUsr {
	return &getLoggedUsr{db}
}

func (g *getLoggedUsr) Run(ctx context.Context, usrID string) (baseuser.BaseUser, error) {
	var user map[string]interface{}
	result := g.db.DB().Table("users").Where("usr_id = ?", usrID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	isStudent := false
	if scope, ok := user["usr_scope"]; ok {
		isStudent = scope == "student"
	}

	if isStudent {
		return student.NewStudent(
			student.WithID(user["usr_id"].(string)),
			student.WithEmail(user["usr_email"].(string)),
			student.WithScope(user["usr_scope"].(string)),
		), nil
	}

	return producer.NewProducer(
		producer.WithID(user["usr_id"].(string)),
		producer.WithEmail(user["usr_email"].(string)),
		producer.WithScope(user["usr_scope"].(string)),
	), nil
}
