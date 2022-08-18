package validator

import "github.com/go-playground/validator/v10"

// ModuleErrorMap is an interface with input structs of data should implements
// to get validated with custom validator.
type ModuleErrorMap interface {
	// GetErrorMap is the method who returns a map with errors.
	GetErrorMap() map[string]map[string]error
}

// CustomValidator is a struct to manage the validator.
type CustomValidator struct {
	validate *validator.Validate
}
