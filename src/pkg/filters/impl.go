package filters

import (
	"fmt"
)

type filters struct {
	fields       []string
	conditionMap ConditionMap
}

// NewConditions will return FilterCondition instance.
func NewConditions() FilterConditions {
	return &filters{
		fields:       make([]string, 0),
		conditionMap: NewConditionMap(),
	}
}

// WithFields will return the fields to be used in the query
// and if has fields to be used.
//
// If WithFields returns false, it means that there are no fields to be used.
func (f *filters) WithFields(prefix string) ([]string, bool) {
	if prefix != "" {
		withPrefix := make([]string, len(f.fields))
		for i, field := range f.fields {
			withPrefix[i] = fmt.Sprintf("%s.%s", prefix, field)
		}

		return withPrefix, len(withPrefix) > 0
	}

	return f.fields, len(f.fields) > 0
}

// HasConditions will return true if there are conditions to be used in the query.
// If HasConditions returns false should NOT use the following methods:
//  - Conditions
//  - GetCondition
func (f *filters) HasConditions() bool {
	return f.conditionMap.Len() > 0
}

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
func (f *filters) Conditions() (string, []interface{}) {
	if !f.HasConditions() {
		panic("No conditions to be used in the query")
	}

	return f.conditionMap.Statement(), f.conditionMap.Values()
}

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
func (f *filters) GetCondition(key string) (interface{}, bool) {
	if !f.HasConditions() {
		panic("No conditions to be used in the query")
	}

	value, ok := f.conditionMap.GetCondition(key)
	return value, ok
}

// WithCondition will add a condition to be used in the query.
func (f *filters) WithCondition(field string, value interface{}) {
	f.conditionMap.AppendCondition(field, EQ, AND, value)
}

// WithComplexCondition will add a complex condition to conditionMap
//
// It can set the ASSERTION type [EQ, NEQ, GT, GTE, etc] and set PREPOSITION
// in Where clause [AND, OR]
//
// Like:
// 	WithComplexCondition("id", EQ, "1", AND)
func (f *filters) WithComplexCondition(field string, assertion Assertion, value interface{}, preposition Preposition) {
	f.conditionMap.AppendCondition(field, assertion, preposition, value)
}

// RemoveCondition will remove a condition to be used in the query.
func (f *filters) RemoveCondition(field string) {
	f.conditionMap.DelCondition(field)
}

// AddField will add a field to be used in the query.
func (f *filters) AddField(field string) {
	f.fields = append(f.fields, field)
}

// AddFields will add a slice of fields to be used in the query.
func (f *filters) AddFields(fields []string) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		f.AddField(field)
	}
}
