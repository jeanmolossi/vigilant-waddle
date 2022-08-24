package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

type getSingleSession struct {
	db database.Database
}

// NewGetSginelSession returns an auth.GetSingleSessionRepository implementation
func NewGetSingleSession(db database.Database) auth.GetSingleSessionRepository {
	return &getSingleSession{
		db: db,
	}
}

// Run will receive filters wich will be used to find the session
// who are requested
func (g *getSingleSession) Run(ctx context.Context, f filters.FilterConditions) (auth.Session, error) {
	if !f.HasConditions() {
		return nil, auth.ErrHasNotSession
	}

	statement, values := f.Conditions()
	session := &SessionModel{}
	result := g.db.DB().Where(statement, values...).First(session)
	if result.Error != nil {
		return nil, result.Error
	}

	return session.AsDomain(), nil
}
