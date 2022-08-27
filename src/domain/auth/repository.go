package auth

import (
	"context"

	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

// CreateSessionRepository is a interface who defines the CreateSession method
type CreateSessionRepository interface {
	// Run the CreateSession method
	//
	// It will receive the following params:
	//
	//		ctx context.Context
	//		session Session
	//
	// The session received will be replaced by current session if exists.
	// If current session is expired or not found, it will create a new session.
	Run(ctx context.Context, session Session) (Session, error)
}

// DeleteSessionRepository is a interface who defines the DeleteSession method
type DeleteSessionRepository interface {
	// Run the DeleteSession method
	//
	// It will receive the following params:
	//
	//		ctx context.Context
	//		sessionID string
	//
	// The sessionID received will be used to delete the session.
	Run(ctx context.Context, sessionID string) error
}

// GetSingleSessionRepository is a interface who defines the GetSingleSession method
type GetSingleSessionRepository interface {
	// Run the GetSingleSession method
	// The sessionID received will be used to get the session.
	Run(ctx context.Context, f filters.FilterConditions) (Session, error)
}

// SessionUpdater is a interface who defines the UpdateSession method
type SessionUpdater func(Session) (Session, error)

type UpdateSessionRepository interface {
	// Run the UpdateSession method
	Run(ctx context.Context, sessionID string, u SessionUpdater) error
}

// GetLoggerUsr will search and retrieve current logged usr
type GetLoggedUsr interface {
	// Run will receive usrID and return the user or an error
	Run(ctx context.Context, f filters.FilterConditions) (baseuser.BaseUser, error)
}
