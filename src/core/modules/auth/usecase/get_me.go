// Package usecase will implements all student usecases
package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

// GetMe is a usecase to get a student
//
// It will return a student or an error
func GetMe(
	ctx context.Context,
	repo auth.GetLoggedUsr,
) auth.GetMe {
	return func(gmo auth.GetMeOptions) (baseuser.BaseUser, error) {
		f := filters.NewConditions()
		f.AddFields(gmo.Fields)
		f.WithCondition("usr_id", gmo.UserID)

		user, err := repo.Run(ctx, f)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
