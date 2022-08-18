package validator

import "fmt"

// ValidationErr is a struct to manage the validation errors.
// Err is the main error message.
// Errors is a slice of errors. That slice will contain the errors of every
// field validation.
type ValidationErr struct {
	Err    string  `json:"error"`
	Errors []error `json:"errors,omitempty"`
}

// AddError is a method to add an error to the slice of errors.
func (v *ValidationErr) AddError(err error) {
	v.Errors = append(v.Errors, err)
}

// Error implements error interface.
func (v *ValidationErr) Error() string {
	return fmt.Sprintf("404: %s", v.Err)
}

// FieldError is a struct to manage the field errors.
type FieldError struct {
	Field string `json:"field" example:"field_name"`
	Err   string `json:"message" example:"field_name is required"`
}

// Error implements error interface.
func (f *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", f.Field, f.Err)
}
