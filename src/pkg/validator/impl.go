package validator

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

// NewCustomValidation is a factory method to create a CustomValidator.
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validate: validator.New(),
	}
}

// Validate is a method to validate a struct with custom validator.
// Param `i` should be an struct tagged with validate tag.
func (cv *CustomValidator) Validate(i interface{}) error {
	// conform.Strings will remove the empty strings from the slice.
	// will parse determined strings as the valid and accepted format.
	conform.Strings(i)

	err := cv.validate.Struct(i)

	// Default bad request validation error
	fieldErrors := &ValidationErr{
		Err:    "Bad Request",
		Errors: []error{},
	}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println("Validation error:", err)
			return err
		}

		// cast our param as ModuleErrorMap
		if errMap, ok := i.(ModuleErrorMap); ok {
			// get the map of errors
			for _, err := range err.(validator.ValidationErrors) {
				if fieldErr := GetFieldError(err, errMap); fieldErr != nil {
					// when we have a field error, add it to the slice of errors
					fieldErrors.AddError(fieldErr)
				}
			}
		} else {
			// If we can't cast the param as ModuleErrorMap, we return the default error
			return errors.New("not module error map")
		}

		return fieldErrors
	}

	return nil
}

// GetFieldError gets a validation field error and ModuleErrorMap.
// That function select the correct field error from the map of errors and
// return that error.
func GetFieldError(err validator.FieldError, moduleErrMap ModuleErrorMap) error {
	errMap := moduleErrMap.GetErrorMap()

	// set the field as lowercase.
	// Example: Title will be title
	lowField := getFieldNamespace(err)
	// Gets error tag. If the error is because required tag, or max tag...
	// Example: required
	errTag := err.Tag()

	// If the field error exists in the map, we return the error
	if errMap[lowField] != nil {
		return &FieldError{
			Field: lowField,
			Err:   errMap[lowField][errTag].Error(),
		}
	}

	// If the field error doesn't exist in the map, we return the default error
	return &FieldError{Field: lowField, Err: "unknown error"}
}

// getFieldNamespace is a function to get the field namespace.
//
// Example: Title will be title
// Example: Title.Subtitle will be title.subtitle
func getFieldNamespace(err validator.FieldError) string {
	lowNamespace := strings.ToLower(err.Namespace())
	propNames := strings.Split(lowNamespace, ".")

	// slice root struct name
	// Example:
	//
	// 		[rootstruct fieldname] => [fieldname]
	//		[rootstruct field fieldname.subfieldname] => [field fieldname.subfieldname]
	rootSpliced := propNames[1:]

	return strings.Join(rootSpliced, ".")
}
