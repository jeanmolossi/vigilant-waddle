package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
)

func EndSession(
	ctx context.Context,
	repo auth.DeleteSessionRepository,
) auth.EndSession {
	return func(sessionID string) error {
		return repo.Run(ctx, sessionID)
	}
}
