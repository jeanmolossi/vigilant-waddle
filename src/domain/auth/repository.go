package auth

import (
	"context"
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
