package producer

import (
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
)

// WithID returns a baseuser.Option that sets the ID of a producer.
//
// It receives a string and sets the ID of the producer.
// If the string is empty, it will return an error.
func WithID(id string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*producer); ok {
			usr.id = id
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}

// WithEmail returns a baseuser.Option that sets the Email of a producer.
//
// It receives a string and sets the Email of the producer.
// If the string is empty, it will return an error.
func WithEmail(email string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*producer); ok {
			usr.email = email
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}

// WithPassword returns a baseuser.Option that sets the Password of a producer.
func WithPassword(password string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*producer); ok {
			usr.password = password
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}

// WithScope will add scope to usr
func WithScope(scope string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*producer); ok {
			usr.scope = baseuser.Scope(scope)
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}
