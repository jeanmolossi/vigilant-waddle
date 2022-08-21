package auth

import (
	"errors"
	"time"
)

var ErrMissingToken = errors.New("missing token")

// WithSessionID sets the session ID for the session
func WithSessionID(id string) SessionProp {
	return func(s *session) error {
		s.sessionID = id
		return nil
	}
}

// WithStudentID sets the student ID for the session
func WithStudentID(id string) SessionProp {
	return func(s *session) error {
		s.studentID = id
		return nil
	}
}

// WithExpiration sets the expiration for the session
func WithExpiration(expiration time.Time) SessionProp {
	return func(s *session) error {
		s.expiration = expiration
		return nil
	}
}

// WithAccessToken sets the access token for the session
func WithAccessToken(token string) SessionProp {
	return func(s *session) error {
		s.accessToken = token
		return nil
	}
}

// WithRefreshToken sets the refresh token for the session
func WithRefreshToken(token string) SessionProp {
	return func(s *session) error {
		s.refreshToken = token
		return nil
	}
}
