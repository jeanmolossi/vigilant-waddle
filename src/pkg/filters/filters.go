// Package filters will be used to filter the results of a query.
//
// That package will contains the FilterConditions interface and
// the filters struct.
//
// The filters struct is private but has the NewFilterConditions factory
// who returns filters instance.
//
// That package was made for easy use with https://gorm.io
package filters

// FilterConditions will be used to filter the results of a query.
type FilterConditions interface {
	// WithFields will return the fields to be used in the query
	// and if has fields to be used.
	//
	// If WithFields returns false, it means that there are no fields to be used.
	WithFields(prefix string) ([]string, bool)
	// HasConditions will return true if there are conditions to be used in the query.
	// If HasConditions returns false should NOT use the following methods:
	//  - Conditions
	//  - GetCondition
	HasConditions() bool
	// Conditions will build the string statement and a values slice to be used in the query.
	//
	// On the following example:
	//
	// 		conditions := NewFilterConditions()
	// 		conditions.WithCondition("course_published", true)
	// 		conditions.WithCondition("course_name", "Effective Eureka")
	//
	// Conditions will return the following:
	//
	// 		statement := "course_published = ? AND course_name = ?"
	// 		values := []interface{}{true, "Effective Eureka"}
	Conditions() (string, []interface{})
	// GetCondition will return the value and if exists a condition.
	//
	// On the following example:
	//
	// 		conditions := NewFilterConditions()
	// 		conditions.WithCondition("course_published", true)
	// 		conditions.WithCondition("course_name", "Effective Eureka")
	//      cond := conditions.GetCondition("course_name")
	//
	// GetCondition will return the following to cond:
	//
	// 		"Effective Eureka", true
	GetCondition(key string) (interface{}, bool)
	// WithCondition will add a condition to be used in the query.
	WithCondition(field string, value interface{})
	// RemoveCondition will remove a condition to be used in the query.
	RemoveCondition(field string)
	// AddField will add a field to be used in the query.
	AddField(field string)
	// AddFields will add a slice of fields to be used in the query.
	AddFields(fields []string)
}
