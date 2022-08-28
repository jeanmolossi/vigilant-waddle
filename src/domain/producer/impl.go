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
func (p *producer) GetID() string { return p.id }

// GetEmail will return the current user email
func (p *producer) GetEmail() string { return p.email }

// GetPassword will return the current user password
func (p *producer) GetPassword() string { return p.password }

// GetScope will return producer scope
func (p *producer) GetScope() baseuser.Scope { return p.scope }

// GetACL will return ACL for the current user
func (p *producer) GetACL() baseuser.ACL { return baseuser.ProducerACL() }

// SyncData receives an array of options and applies them to the current user
func (p *producer) SyncData(usrOptions ...baseuser.Option) error {
	if len(usrOptions) == 0 {
		return ErrNoDataToSync
	}

	for _, opt := range usrOptions {
		if err := opt(p); err != nil {
			return err
		}
	}

	return nil
}

// IsPasswordValid will check if the user password received
// is valid password for the current user
func (p *producer) IsPasswordValid(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.password), []byte(password)) == nil
}

// HashPassword will hash the password received and return the hash
//
// It should have a password, else it will return an error
func (p *producer) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.password = string(hash)
	return nil
}

// AddScope will add current scope to current user.
//
// Looks possible cases:
//
// 		student.AddScope(PRODUCER)	// student.GetScope()	=> BOTH
// 		producer.AddScope(STUDENT)	// producer.GetScope()	=> BOTH
// 		student.AddScope(STUDENT) 	// student.GetScope()	=> STUDENT
// 		producer.AddScope(PRODUCER)	// producer.GetScope()	=> PRODUCER
// 		student.AddScope(BOTH) 		// student.GetScope()	=> BOTH
// 		producer.AddScope(BOTH)		// producer.GetScope()	=> BOTH
//
// If user scope is `student` and call `AddScope(PRODUCER)`, it
// will make user scope as `BOTH`.
//
// If user scope is `producer` and call `AddScope(STUDENT)`, it
// will make user scope as `BOTH` too.
//
// If current scope matches with call `AddScope`. Nothing happens
func (p *producer) AddScope(scope baseuser.Scope) {
	if p.scope == baseuser.STUDENT {
		p.scope = baseuser.BOTH
	}
}
