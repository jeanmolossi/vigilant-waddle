package baseuser

const (
	GetMe  Resource = "get_me"
	Login  Resource = "login"
	Logout Resource = "logout"
)

// StudentACL is the ACL for a student.
//
// The ACL preconfigured for a student.
func StudentACL() ACL {
	return NewACL(
		// ReadCapabilities
		CanRead(GetMe),
		// FullAccessCapabilities
		FullAccess(Login),
		FullAccess(Logout),
	)
}

// ProducerACL is the ACL for a producer.
//
// The ACL preconfigured for a producer.
func ProducerACL() ACL {
	return NewACL(
		// ReadCapabilities
		CanRead(GetMe),
		// FullAccessCapabilities
		FullAccess(Login),
		FullAccess(Logout),
	)
}
