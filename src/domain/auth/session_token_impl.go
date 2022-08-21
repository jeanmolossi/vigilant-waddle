package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// sessionToken is a struct who implements SessionToken interface.
type sessionToken struct {
	token      string
	expiration time.Time
	issuer     string
}

var (
	SECRET            = []byte(os.Getenv("SESSION_SECRET"))
	AccessExpiration  = time.Minute * 10   // 10 minutes
	RefreshExpiration = time.Hour * 24 * 7 // 1 week

	AccessIssuer  = "session_provider"
	RefreshIssuer = "refresh_provider"
)

// SessionFromToken will build a SessionToken from a token string.
func SessionFromToken(token string) SessionToken {
	decodedToken, err := decodeSessionToken(token)
	if err != nil {
		// Return a non nil sessionToken to avoid panic.
		// The session is expired or invalid.
		return NewSessionToken("")
	}

	return decodedToken
}

// NewSessionToken will build a new SessionToken.
//
// It receives required studentID param and a list of options.
//
// Hint:
// Use WithIssuer() to change the type of token. (default is "session_provider")
// Use ExpiresIn() to set the expiration for the token. (default is 10 minutes)
func NewSessionToken(studentID string, opts ...TokenOption) SessionToken {
	s := &sessionToken{
		issuer:     AccessIssuer,
		expiration: time.Now().UTC().Local().Add(AccessExpiration),
	}

	if len(opts) > 0 {
		for _, applyOpt := range opts {
			if err := applyOpt(s); err != nil {
				panic(err.Error())
			}
		}
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: s.expiration.Unix(),
		Subject:   studentID,
		Issuer:    s.issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, _ := token.SignedString(SECRET)

	return &sessionToken{
		token:      signedTokenString,
		expiration: s.expiration,
	}
}

// Token returns the token string.
func (s *sessionToken) Token() string {
	return s.token
}

// Expiration returns the expiration time.
func (s *sessionToken) Expiration() time.Time {
	return s.expiration
}

// IsExpired returns true if the token is expired.
func (s *sessionToken) IsExpired() bool {
	return s.expiration.Before(time.Now().UTC().Local())
}

// decodeSessionToken will receive token string and decode it.
func decodeSessionToken(token string) (SessionToken, error) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	return &sessionToken{
		token:      token,
		expiration: time.Unix(claims.ExpiresAt, 0),
	}, nil
}
