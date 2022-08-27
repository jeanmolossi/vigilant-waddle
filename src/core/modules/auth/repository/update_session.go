package repository

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type updateSession struct {
	db database.Database
}

// NewUpdateSession returns an auth.UpdateSessionRepository implementation
func NewUpdateSession(db database.Database) auth.UpdateSessionRepository {
	return &updateSession{
		db: db,
	}
}

// Run will update given sessionID by the update call
//
// Example:
//
// 	// current session will be updated
// 	updater := func(cs auth.Session) (auth.Session, error) {
// 		updatedSession := auth.NewSession(
// 			auth.WithSessionID(cs.GetID()),
// 			auth.WithSessionID(cs.GetStudentID()),
// 		)
// 	}
//
// 	repo := updateSessionRepository.Run(
// 		context.Background(),
// 		sessionID,
// 		updater,
// 	)
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
		UserID:       session.GetUserID(),
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
