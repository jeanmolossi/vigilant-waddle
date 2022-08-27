package filters

import (
	"fmt"
	"strings"
)

// condition will store configuration to
// map a single condition
type condition struct {
	index       int
	Field       string
	Assertion   Assertion
	Preposition Preposition
	values      []interface{}
}

func NewCondition(
	index int,
	f string,
	a Assertion,
	p Preposition,
	values ...interface{},
) Condition {
	return &condition{
		index,
		f, a, p,
		values,
	}
}

// Statement mount the string statement to return condition
func (c *condition) Statement(isFirst bool) string {
	statement := ""
	format := "%s%s %s %s"

	// when is NOT the first statement in loop
	// we add preposition in statement
	if !isFirst {
		format = "%s %s %s %s"
		statement = string(c.Preposition)
	}

	// when the assertion can receive multiple values
	// Looks like:
	// 	IN (?, ?)
	if assertionIncludes(multipleValuesAssertions(), c.Assertion) {
		format = "%s %s %s (%s)"
	}

	statement = fmt.Sprintf(
		format,
		statement,
		c.Field,
		c.Assertion,
		strings.Repeat("?", len(c.values)),
	)

	return statement
}

// Value will return the received value to condition
// 	It will return nil if has no values received.
// 	It will return a single interface when is single value.
// 	It will return a slice of interfaces when has more than one
// value received
func (c *condition) Value() interface{} {
	if len(c.values) == 0 {
		return nil
	}

	if len(c.values) == 1 {
		return c.values[0]
	}

	return c.values
}

func multipleValuesAssertions() []Assertion {
	return []Assertion{IN}
}
