package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
	"gorm.io/gorm"
)

type deleteSession struct {
	db database.Database
}

// NewDeleteSession will return auth.DeleteSessionRepository
func NewDeleteSession(db database.Database) auth.DeleteSessionRepository {
	return &deleteSession{db}
}

// Run will delete received sessionID.
// Useful to invalidate a session
func (d *deleteSession) Run(ctx context.Context, sessionID string) error {
	s := &SessionModel{SessID: sessionID}

	result := d.db.DB().First(s)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}

		return result.Error
	}

	return d.db.DB().Delete(s).Error
}
