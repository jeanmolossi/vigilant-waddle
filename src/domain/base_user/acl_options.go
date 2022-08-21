package baseuser

// CanRead set ACLOptions to allow read access to the given resource.
func CanRead(resource Resource) ACLOptions {
	return func(a *acl) {
		a.canRead[resource] = true
	}
}

// CanWrite set ACLOptions to allow write access to the given resource.
func CanWrite(resource Resource) ACLOptions {
	return func(a *acl) {
		a.canWrite[resource] = true
	}
}

// FullAccess set ACLOptions to allow read and write access to the given resource.
func FullAccess(resource Resource) ACLOptions {
	return func(a *acl) {
		a.canRead[resource] = true
		a.canWrite[resource] = true
	}
}
