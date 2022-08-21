package repository

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SessionToken struct {
	Token      string
	Expiration time.Time
}

var (
	SECRET = []byte(os.Getenv("SESSION_SECRET"))
)

func NewSessionToken(studentID string) *SessionToken {
	expiration := time.Now().UTC().Local().Add(time.Minute * 10)

	claims := &jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
		Subject:   studentID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, _ := token.SignedString(SECRET)

	return &SessionToken{
		Token:      signedTokenString,
		Expiration: expiration,
	}
}

func DecodeSessionToken(token string) (*SessionToken, error) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	return &SessionToken{
		Token:      token,
		Expiration: time.Unix(claims.ExpiresAt, 0),
	}, nil
}
