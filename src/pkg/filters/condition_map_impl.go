package filters

import "strings"

// conditionMap will store the conditions to
// mount where clouse in query
type conditionMap struct {
	conditions map[string]Condition

	preparedValues    []interface{}
	preparedStatement string
}

// NewConditionMap will return ConditionMap implementation
func NewConditionMap() ConditionMap {
	return &conditionMap{
		make(map[string]Condition, 0),
		nil,
		"",
	}
}

// Values will prepare statement to return and will return values
func (c *conditionMap) Values() []interface{} {
	if c.preparedValues != nil {
		return c.preparedValues
	}

	values := make([]interface{}, 0)
	statement := []string{}

	counter := 0
	for _, condition := range c.conditions {
		counter++

		statement = append(statement,
			condition.Statement(
				counter == 1,
			),
		)
		values = append(values, condition.Value())
	}

	c.preparedStatement = strings.Join(statement, " ")
	c.preparedValues = values
	return values
}

// Statement will prepare statement to return and will return
// string statement
func (c *conditionMap) Statement() string {
	if c.preparedStatement != "" {
		return c.preparedStatement
	}

	values := make([]interface{}, 0)
	statement := []string{}

	counter := 0
	for _, condition := range c.conditions {
		counter++

		statement = append(statement,
			condition.Statement(
				counter == 1,
			),
		)
		values = append(values, condition.Value())
	}

	c.preparedStatement = strings.Join(statement, " ")
	c.preparedValues = values
	return strings.Join(statement, " ")
}

// Len will return how many condition are mapped
func (c *conditionMap) Len() int {
	return len(c.conditions)
}

// AppendCondition will add a condition to map
func (c *conditionMap) AppendCondition(field string, assertion Assertion, preposition Preposition, values ...interface{}) {
	condition := NewCondition((c.Len()), field, assertion, preposition, values...)
	c.conditions[field] = condition
}

// DelCondition remove mapped condition
func (c *conditionMap) DelCondition(field string) {
	delete(c.conditions, field)
}

// GetCondition will return a single condition found by field
func (c *conditionMap) GetCondition(field string) (interface{}, bool) {
	if c.Len() == 0 {
		return nil, false
	}

	value, ok := c.conditions[field]
	return value, ok
}
