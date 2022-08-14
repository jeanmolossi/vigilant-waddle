package filters_test

import (
	"fmt"

	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
)

func Example() {
	f := filters.NewConditions()
	f.AddFields([]string{"course_published", "course_name"})
	f.WithCondition("course_published", true)
	f.WithCondition("course_name", "Effective Eureka")

	fields, hasFields := f.WithFields("prefixed")
	statement, values := f.Conditions()
	fmt.Println(fields, hasFields)
	fmt.Println(statement)
	fmt.Println(values)

	// Output:
	// [prefixed.course_published prefixed.course_name] true
	// course_published = ? AND course_name = ?
	// [true Effective Eureka]
}
