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
	// GetScope will return the user scope
	GetScope() Scope
	// GetACL will return ACL for the current user
	GetACL() ACL

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

// Resource is a custom type to represent the different resources
// that can be used in the ACL
type Resource string

// String will only parse resource to a string
func (r Resource) String() string {
	return string(r)
}

// ACLOptions is a function that will be used to configure the ACL
type ACLOptions func(a *acl)

// ACL will provide the access control list for the current user
// and the resources that can be accessed by the current user
type ACL interface {
	// CanRead will check if the current user can read the resource
	CanRead(resource Resource) bool
	// CanWrite will check if the current user can write the resource
	CanWrite(resource Resource) bool
	// FullAccess will check if the current user can read and write the resource
	FullAccess(resource Resource) bool
	// GetResources will return two slices from resources.
	// Readable resources and Writable resources in order.
	GetResources() ([]Resource, []Resource)
}
