package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
	"gorm.io/gorm"
)

type createSession struct {
	db database.Database
}

// NewCreateSession return a auth.CreateSessionRepository implementation
func NewCreateSession(db database.Database) auth.CreateSessionRepository {
	return &createSession{db}
}

// Run will try to create the incoming session
//
// If already exists a session, will return that.
//
// Else, if session not exists create one.
func (c *createSession) Run(ctx context.Context, session auth.Session) (auth.Session, error) {
	if session.GetUserID() == "" {
		return nil, auth.ErrHasNotStudentID
	}

	actualSession := &SessionModel{
		UserID: session.GetUserID(),
	}

	result := c.db.DB().Where("usr_id = ?", session.GetUserID()).First(actualSession)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		}
	}

	if result.RowsAffected > 0 {
		// In cases wich session is not expired, return that
		if !actualSession.AsDomain().IsExpired() {
			return actualSession.AsDomain(), nil
		} else {
			// Delete current session if expired
			c.db.DB().Delete(actualSession)
		}
	}

	// updates current session with incoming session
	// Incoming session ever is valid
	actualSession.AccessToken = session.GetAccessToken().Token()
	actualSession.RefreshToken = session.GetRefreshToken().Token()
	actualSession.Expiration = session.GetAccessToken().Expiration()

	// create the updated session in db
	result = c.db.DB().Create(actualSession)
	if result.Error != nil {
		return nil, result.Error
	}

	return actualSession.AsDomain(), nil
}
