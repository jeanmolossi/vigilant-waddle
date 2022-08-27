package auth

import (
	"encoding/base64"
	"fmt"
	"time"
)

// session implements the Session interface.
type session struct {
	sessionID    string
	userID       string
	expiration   time.Time
	accessToken  string
	refreshToken string
}

// NewSession will create a new Session instance.
//
// It creates a new Valid session if only if receives the following options:
//
//		s := auth.NewSession(
//			auth.WithSessionID(sessionID),
//			auth.WithUserID(userID),
//			auth.WithExpiration(expiration),
//			auth.WithAccessToken(accessToken),
//			auth.WithRefreshToken(refreshToken),
//		)
//
// If any of the options is missing can return a invalid or expired session.
func NewSession(opts ...SessionProp) Session {
	// this will initialize the session with default values
	// if no options are provided our session is empty
	s := &session{
		sessionID: "",
		userID:    "",
		// expiration as 10 seconds before now.
		// this is to make sure that the session is expired before it is created as default.
		expiration: time.Now().Add(time.Second * -10),
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if err := opt(s); err != nil {
				panic(err.Error())
			}
		}
	}

	return s
}

// GetID returns the session ID.
func (s *session) GetID() string { return s.sessionID }

// GetUserID returns the student ID.
func (s *session) GetUserID() string { return s.userID }

// IsExpired returns true if the session is expired.
//
// If the session has no one of:
//		access token
//		refresh token
//	or
//		expiration
// it will return true to IsExpired call.
//
// Other case is:
//
// Access token or refresh token is expired.
func (s *session) IsExpired() bool {
	if s.accessToken == "" || s.refreshToken == "" {
		return true
	}

	accessToken := s.GetAccessToken()
	refreshToken := s.GetRefreshToken()

	if accessToken.IsExpired() || refreshToken.IsExpired() {
		return true
	}

	if s.expiration.IsZero() {
		return true
	}

	return s.expiration.Before(time.Now())
}

// GetAccessToken returns the access token.
func (s *session) GetAccessToken() SessionToken {
	return SessionFromToken(s.accessToken)
}

// GetRefreshToken returns the refresh token.
func (s *session) GetRefreshToken() SessionToken {
	return SessionFromToken(s.refreshToken)
}

// HashToken will hash the token and return the hash.
//
// The hash is used to return the token to the client.
func (s *session) HashToken() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", s.userID, s.sessionID)),
	)
}
