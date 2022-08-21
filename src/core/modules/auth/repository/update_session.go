package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type updateSession struct {
	db database.Database
}

func NewUpdateSession(db database.Database) auth.UpdateSessionRepository {
	return &updateSession{
		db: db,
	}
}

func (u *updateSession) Run(ctx context.Context, sessionID string, update auth.SessionUpdater) error {
	currentSession := &SessionModel{SessID: sessionID}

	result := u.db.DB().First(currentSession)
	if result.Error != nil {
		return result.Error
	}

	session, err := update(currentSession.AsDomain())
	if err != nil {
		return err
	}

	updatedSession := &SessionModel{
		SessID:       session.GetID(),
		StudentID:    session.GetStudentID(),
		AccessToken:  session.GetAccessToken().Token(),
		RefreshToken: session.GetRefreshToken().Token(),
		Expiration:   session.GetAccessToken().Expiration(),
	}

	result = u.db.DB().Save(updatedSession)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
