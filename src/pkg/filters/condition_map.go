package filters

// ConditionMap will handle all conditions and how it will
// be compiled and formated.
type ConditionMap interface {
	// Len returns counting of how many conditions are set
	Len() int
	// Values will result in a slice of received condition values
	Values() []interface{}
	// Statement returns formatted statment string.
	//
	// Like:
	// 	`id = ? AND name = ? OR ip IN (?)`
	Statement() string
	// AppendCondition will append a condition with received config inside
	// ConditionMap
	//
	// Usage:
	// 	AppendCondition("id", filters.IN, filters.AND, []string{"1","2"})
	// Statment will be:
	// 	`id IN (?)`
	// Values will be
	// 	[]interface{"1", "2"}
	AppendCondition(field string, assertion Assertion, preposition Preposition, values ...interface{})
	// DelCondition will remove a field defined condition
	DelCondition(field string)
	// GetCondition will get a field condition configuration
	//
	// If condition does not set will return false on the second returned value
	GetCondition(field string) (interface{}, bool)
}

// Condition is the interface to handle a single Condition inside
// the ConditionMap
type Condition interface {
	// Statement build statement string to ConditionMap
	//
	// It receives "isFirst" param to define if it
	// will receive preposition on statement.
	//
	// In cases who isFirst becomes true, the preposition
	// will be ignored in final statment
	Statement(isFirst bool) string
	// Value will handle received value to return
	// a single value if is a single value defined
	// in the Condition configuration
	//
	// To slices or multiple values it will handle
	// to return value as received
	Value() interface{}
}
