package baseuser

// Scope is a config option to users.
//
// It determines if user is a student, producer or both
type Scope string

const (
	STUDENT  Scope = "student"
	PRODUCER Scope = "producer"
	BOTH     Scope = "both"
)

// String will return scope as string
func (s Scope) String() string {
	return string(s)
}

// Check if the user implementation is Student
func IsStudent(b BaseUser) bool {
	return b.GetScope() == STUDENT || b.GetScope() == BOTH
}

// Check if the user implementation is producer
func IsProducer(b BaseUser) bool {
	return b.GetScope() == PRODUCER || b.GetScope() == BOTH
}
