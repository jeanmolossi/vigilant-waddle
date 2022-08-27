package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

func ValidateAndRefresh(
	ctx context.Context,
	getSessionRepo auth.GetSingleSessionRepository,
	refreshRepo auth.UpdateSessionRepository,
	getLoggedRepo auth.GetLoggedUsr,
) auth.ValidateAndRefresh {
	getSession := func(usrID, sessionID string) (auth.Session, error) {
		f := filters.NewConditions()
		f.WithCondition("usr_id", usrID)
		f.WithCondition("session_id", sessionID)

		return getSessionRepo.Run(ctx, f)
	}

	refreshSession := func(old auth.Session) auth.SessionUpdater {
		return func(s auth.Session) (auth.Session, error) {
			isAcessExpired := old.GetAccessToken().IsExpired()
			isRefreshExpired := old.GetRefreshToken().IsExpired()

			accessToken := s.GetAccessToken()
			if isAcessExpired {
				accessToken = auth.NewSessionToken(s.GetUserID())
			}

			refreshToken := s.GetRefreshToken()
			if isRefreshExpired {
				refreshToken = auth.NewSessionToken(
					s.GetUserID(),
					auth.ExpiresIn(auth.RefreshExpiration),
					auth.WithIssuer(auth.RefreshIssuer),
				)
			}

			s = auth.NewSession(
				auth.WithSessionID(s.GetID()),
				auth.WithUserID(s.GetUserID()),
				auth.WithAccessToken(accessToken.Token()),
				auth.WithRefreshToken(refreshToken.Token()),
				auth.WithExpiration(accessToken.Expiration()),
			)

			return s, nil
		}
	}

	return func(usrID, sessionID string) (baseuser.BaseUser, error) {
		session, err := getSession(usrID, sessionID)
		if err != nil {
			return nil, err
		}

		f := filters.NewConditions()
		f.WithCondition("usr_id", usrID)

		// is current session is valid continue
		if !session.IsExpired() {
			return getLoggedRepo.Run(ctx, f)
		}

		if err := refreshRepo.Run(ctx, sessionID, refreshSession(session)); err != nil {
			return nil, err
		}

		return getLoggedRepo.Run(ctx, f)
	}
}
