package baseuser

const (
	GetMe Resource = "get_me"
	Login Resource = "login"
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
	)
}
