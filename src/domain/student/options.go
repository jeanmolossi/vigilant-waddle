package student

import (
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
)

// WithID returns a baseuser.Option that sets the ID of a student.
//
// It receives a string and sets the ID of the student.
// If the string is empty, it will return an error.
func WithID(id string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*student); ok {
			usr.id = id
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}

// WithEmail returns a baseuser.Option that sets the Email of a student.
//
// It receives a string and sets the Email of the student.
// If the string is empty, it will return an error.
func WithEmail(email string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*student); ok {
			usr.email = email
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}

// WithPassword returns a baseuser.Option that sets the Password of a student.
func WithPassword(password string) baseuser.Option {
	return func(u baseuser.BaseUser) error {
		if usr, ok := u.(*student); ok {
			usr.password = password
			return nil
		}

		return baseuser.ErrInvalidBaseUsrImplementation
	}
}
