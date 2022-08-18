package validator_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jeanmolossi/vigilant-waddle/src/pkg/validator"
)

// InputToValidate is a struct to validate example.
//
// Validation tags are based on the validator package.
// Read more at:
// https://pkg.go.dev/github.com/go-playground/validator/v10?utm_source=godoc#pkg-overview
type InputToValidate struct {
	Field  string `validate:"required"`
	Nested struct {
		Field string `validate:"required"`
	} `validate:"required"`
}

// GetErrorMap implements validator.ModuleErrorMap interface.
func (i *InputToValidate) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"field": {
			"required": errors.New("field is required"),
		},
		"nested": {
			"required": errors.New("nested is required"),
		},
		"nested.field": {
			"required": errors.New("nested field is required"),
		},
	}
}

func Example() {
	i := &InputToValidate{}
	v := validator.NewCustomValidator()
	err := v.Validate(i)

	result := errToInterface(err)

	fmt.Println(result)
	// Output:
	// 	map[error:Bad Request errors:[map[field:field message:field is required] map[field:nested.field message:nested field is required]]]
}

// errToInterface converts error to interface.
//
// This is a helper function to convert error to interface.
// This is useful to print error in a pretty way.
func errToInterface(err error) interface{} {
	var result interface{}
	bytes, _ := json.Marshal(err)
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Fatal(err)
	}

	return result
}
