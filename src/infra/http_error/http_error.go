package http_error

import "github.com/jeanmolossi/vigilant-waddle/src/pkg/validator"

type HTTPError struct {
	Err string `json:"error" example:"error message"`
}

type HTTPBadRequestError struct {
	Err  string                 `json:"error" example:"error message"`
	Errs []validator.FieldError `json:"errors"`
}
