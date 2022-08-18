// Package baseuser is a domain package wich contains the common users features
// and behavior.
package baseuser

// Option is a optional function argument for BaseUser
//
// It will be used to configure the BaseUser instance
type Option func(u BaseUser) error

type BaseUser interface {
	// GetID will return the current user id
	GetID() string
	// GetEmail will return the current user email
	GetEmail() string
	// GetPassword will return the current user password
	GetPassword() string

	// SyncData receives an array of options and applies them to the current user
	SyncData(usrOption ...Option) error

	// IsPasswordValid will check if the user password received
	// is valid password for the current user
	IsPasswordValid(password string) bool

	// HashPassword will hash the password received and return the hash
	//
	// It should have a password, else it will return an error
	HashPassword() error
}
