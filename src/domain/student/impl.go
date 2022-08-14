package student

import (
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
)

// student implements Student interface wich contains the BaseUser interface
// student will be used to store the student data
type student struct {
	id       string
	email    string
	password string
}

// NewStudent returns a new student instance.
//
// It receive an array of options and applies them to the student instance.
// If has no options, the instance will be empty.
func NewStudent(usrOptions ...baseuser.Option) Student {
	s := &student{}
	if len(usrOptions) > 0 {
		s.SyncData(usrOptions...)
	}
	return s
}

// GetID will return the current user id
func (s *student) GetID() string { return s.id }

// GetEmail will return the current user email
func (s *student) GetEmail() string { return s.email }

// SyncData receives an array of options and applies them to the current user
func (s *student) SyncData(usrOptions ...baseuser.Option) error {
	if len(usrOptions) == 0 {
		return ErrNoDataToSync
	}

	for _, opt := range usrOptions {
		if err := opt(s); err != nil {
			return err
		}
	}

	return nil
}

// IsPasswordValid will check if the user password received
// is valid password for the current user
func (s *student) IsPasswordValid(password string) bool { panic("not implemented") }

// HashPassword will hash the password received and return the hash
//
// It should have a password, else it will return an error
func (s *student) HashPassword() error { panic("not implemented") }
