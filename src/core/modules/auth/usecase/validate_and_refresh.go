package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

func ValidateAndRefresh(
	ctx context.Context,
	getSessionRepo auth.GetSingleSessionRepository,
	refreshRepo auth.UpdateSessionRepository,
) auth.ValidateAndRefresh {
	getSession := func(studentID, sessionID string) (auth.Session, error) {
		f := filters.NewConditions()
		f.WithCondition("student_id", studentID)
		f.WithCondition("session_id", sessionID)

		return getSessionRepo.Run(ctx, f)
	}

	refreshSession := func(old auth.Session) auth.SessionUpdater {
		return func(s auth.Session) (auth.Session, error) {
			isAcessExpired := old.GetAccessToken().IsExpired()
			isRefreshExpired := old.GetRefreshToken().IsExpired()

			accessToken := s.GetAccessToken()
			if isAcessExpired {
				accessToken = auth.NewSessionToken(s.GetStudentID())
			}

			refreshToken := s.GetRefreshToken()
			if isRefreshExpired {
				refreshToken = auth.NewSessionToken(
					s.GetStudentID(),
					auth.ExpiresIn(auth.RefreshExpiration),
					auth.WithIssuer(auth.RefreshIssuer),
				)
			}

			s = auth.NewSession(
				auth.WithSessionID(s.GetID()),
				auth.WithStudentID(s.GetStudentID()),
				auth.WithAccessToken(accessToken.Token()),
				auth.WithRefreshToken(refreshToken.Token()),
				auth.WithExpiration(accessToken.Expiration()),
			)

			return s, nil
		}
	}

	return func(studentID, sessionID string) error {
		session, err := getSession(studentID, sessionID)
		if err != nil {
			return err
		}

		// is current session is valid continue
		if !session.IsExpired() {
			return nil
		}

		return refreshRepo.Run(ctx, sessionID, refreshSession(session))
	}
}
