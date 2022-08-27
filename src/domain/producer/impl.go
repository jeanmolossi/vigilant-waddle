package producer

import (
	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"golang.org/x/crypto/bcrypt"
)

// producer implements Producer interface wich contains the BaseUser interface
// producer will be used to store the producer data
type producer struct {
	id       string
	email    string
	password string
	scope    baseuser.Scope
}

// NewProducer returns a new producer instance.
//
// It receive an array of options and applies them to the producer instance.
// If has no options, the instance will be empty.
func NewProducer(usrOptions ...baseuser.Option) Producer {
	s := &producer{}

	if len(usrOptions) > 0 {
		err := s.SyncData(usrOptions...)
		if err != nil {
			panic(err.Error())
		}
	}

	return s
}

// GetID will return the current user id
func (s *producer) GetID() string { return s.id }

// GetEmail will return the current user email
func (s *producer) GetEmail() string { return s.email }

// GetPassword will return the current user password
func (s *producer) GetPassword() string { return s.password }

// GetScope will return producer scope
func (s *producer) GetScope() baseuser.Scope { return s.scope }

// GetACL will return ACL for the current user
func (s *producer) GetACL() baseuser.ACL { return baseuser.ProducerACL() }

// SyncData receives an array of options and applies them to the current user
func (s *producer) SyncData(usrOptions ...baseuser.Option) error {
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
func (s *producer) IsPasswordValid(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(s.password), []byte(password)) == nil
}

// HashPassword will hash the password received and return the hash
//
// It should have a password, else it will return an error
func (s *producer) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(s.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.password = string(hash)
	return nil
}
