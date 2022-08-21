package baseuser

type acl struct {
	canRead  map[Resource]bool
	canWrite map[Resource]bool
}

// NewACL returns a new ACL with permissions configured by the given options.
func NewACL(opts ...ACLOptions) ACL {
	if len(opts) == 0 {
		panic("ACL must have at least one option")
	}

	// this will initialize the ACL with not nil values
	a := &acl{
		canRead:  make(map[Resource]bool),
		canWrite: make(map[Resource]bool),
	}

	// apply the options
	for _, opt := range opts {
		opt(a)
	}

	return a
}

// CanRead returns true if the given resource can be read.
func (a *acl) CanRead(resource Resource) bool {
	if canRead, ok := a.canRead[resource]; ok {
		return canRead
	}

	return false
}

// CanWrite returns true if the given resource can be written.
func (a *acl) CanWrite(resource Resource) bool {
	if canWrite, ok := a.canWrite[resource]; ok {
		return canWrite
	}

	return false
}

// FullAccess returns true if the given resource can be read and written.
func (a *acl) FullAccess(resource Resource) bool {
	return a.CanRead(resource) && a.CanWrite(resource)
}
