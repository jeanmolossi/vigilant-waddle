package validator_test

import (
	"errors"
	"testing"

	v "github.com/go-playground/validator/v10"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/validator"
	"github.com/stretchr/testify/assert"
)

type validationErr struct {
	Field string `validate:"required"`
}

func (v *validationErr) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"field": {
			"required": errors.New("field is required"),
		},
	}
}

type validationErrWithNested struct {
	Field  string `validate:"required"`
	Nested struct {
		Field string `validate:"required"`
	} `validate:"required"`
}

func (v *validationErrWithNested) GetErrorMap() map[string]map[string]error {
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

func TestGetFieldError(t *testing.T) {
	t.Run("should return correctly err", func(t *testing.T) {
		i := &validationErr{}
		err := v.New().Struct(i)

		fieldErrs := err.(v.ValidationErrors)
		fieldErr := validator.GetFieldError(fieldErrs[0], i)

		assert.NotNil(t, fieldErr)
		assert.Equal(t, "field: field is required", fieldErr.Error())
	})

	t.Run("should validate nested err", func(t *testing.T) {
		i := &validationErrWithNested{}
		err := v.New().Struct(i)

		fieldErrs := err.(v.ValidationErrors)

		fieldErr := validator.GetFieldError(fieldErrs[0], i)

		var validatorFieldErr *validator.FieldError

		assert.NotNil(t, fieldErr)
		assert.EqualError(t, fieldErr, "field: field is required")
		assert.ErrorAs(t, fieldErr, &validatorFieldErr)
	})
}
