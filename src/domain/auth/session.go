// Package auth will contains all the logic related to the authentication of the user.
package auth

import "time"

// SessionProp is a function that will be applied to the session.
//
// It will be used to create config options to Session.
type SessionProp func(*session) error

// TokenOption is a function that will be applied to the session token.
//
// It will be used to create config options to SessionToken.
type TokenOption func(*sessionToken) error

// Session is the interface that represents a session.
//
// It will be used to store the authentication information of the user.
type Session interface {
	// GetID returns the session ID.
	GetID() string
	// GetStudentID returns the student ID.
	GetStudentID() string
	// IsExpired returns true if the session is expired.
	IsExpired() bool
	// GetAccessToken returns the access token.
	GetAccessToken() SessionToken
	// GetRefreshToken returns the refresh token.
	GetRefreshToken() SessionToken
	// HashToken returns the hash of the token.
	//
	// It encodes the token format and returns the encoded string.
	// Useful to return the token to the client.
	HashToken() string
}

// SessionToken is the interface that represents a session token.
//
// It will be used to store the authentication token.
type SessionToken interface {
	// Token will return token string
	Token() string
	// Expiration will return the expiration time of the token.
	Expiration() time.Time
	// IsExpired will return true if the token is expired.
	IsExpired() bool
}

// Usecases

// StartSession is the function that will be used to start a session.
type StartSession func(username, password string) (Session, error)
