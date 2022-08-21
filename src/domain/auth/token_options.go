package auth

import "time"

// ExpiresIn will set the expiration for the token.
func ExpiresIn(expiresIn time.Duration) TokenOption {
	return func(st *sessionToken) error {
		st.expiration = time.Now().UTC().Local().Add(expiresIn)
		return nil
	}
}

// WithIssuer sets the issuer for the token.
func WithIssuer(issuer string) TokenOption {
	return func(st *sessionToken) error {
		st.issuer = issuer
		return nil
	}
}
