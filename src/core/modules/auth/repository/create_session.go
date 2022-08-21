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

func NewCreateSession(db database.Database) auth.CreateSessionRepository {
	return &createSession{db}
}

func (c *createSession) Run(ctx context.Context, session auth.Session) (auth.Session, error) {
	if session.GetStudentID() == "" {
		return nil, auth.ErrHasNotStudentID
	}

	actualSession := &SessionModel{
		StudentID: session.GetStudentID(),
	}

	result := c.db.DB().Where("student_id = ?", session.GetStudentID()).First(actualSession)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		}
	}

	if result.RowsAffected > 0 {
		if !actualSession.AsDomain().IsExpired() {
			return actualSession.AsDomain(), nil
		} else {
			c.db.DB().Delete(actualSession)
		}
	}

	actualSession.AccessToken = session.GetAccessToken().Token()
	actualSession.RefreshToken = session.GetRefreshToken().Token()
	actualSession.Expiration = session.GetAccessToken().Expiration()

	result = c.db.DB().Create(actualSession)
	if result.Error != nil {
		return nil, result.Error
	}

	return actualSession.AsDomain(), nil
}
