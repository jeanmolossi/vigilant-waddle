// Package auth will contains all the logic related to the authentication of the user.
package auth

import (
	"errors"
	"time"

	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
)

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
	// GetUserID returns the user ID.
	GetUserID() string
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

// EndSession is the function that will be used to end a session.
type EndSession func(sessionID string) error

// ValidateAndRefresh is the function that will be used to validate and refresh a session.
type ValidateAndRefresh func(usrID, sessionID string) (baseuser.BaseUser, error)

// GetMe is the usecase to get the current student
// It will use the session token
type GetMe func(GetMeOptions) (baseuser.BaseUser, error)

// GetMeOptions is the options for the GetMe usecase
type GetMeOptions struct {
	UserID string   `json:"user_id" example:"1" format:"uuid" validate:"required,uuid"`
	Fields []string `query:"fields" example:"name,email"`
}

// GetErrorMap implements ModuleErrorMap to GetMeOptions
func (g *GetMeOptions) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"userid": {
			"required": errors.New(""),
			"uuid":     errors.New("")},
	}
}
