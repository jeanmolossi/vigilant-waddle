package http_error

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/validator"
	"github.com/labstack/echo/v4"
)

func Handle(c echo.Context, err error) error {
	var validationErr *validator.ValidationErr
	var bindErr *echo.HTTPError

	switch {
	case errors.Is(err, student.ErrEmailAlreadyExists):
		return c.JSON(http.StatusConflict, ToJsonErr(err))

	case errors.As(err, &validationErr),
		errors.As(err, &bindErr):
		return c.JSON(http.StatusBadRequest, err)

	case errors.Is(err, auth.ErrHasNotSession),
		errors.Is(err, auth.ErrHasNotStudentID):
		return c.JSON(http.StatusBadRequest, ToJsonErr(err))

	case errors.Is(err, student.ErrInvalidCredentials),
		errors.Is(err, student.ErrMissingStudentID):
		return c.JSON(http.StatusUnauthorized, ToJsonErr(err))

	case errors.Is(err, auth.ErrForbidden):
		return c.JSON(http.StatusForbidden, ToJsonErr(err))

	default:
		return c.JSON(http.StatusInternalServerError, ToJsonErr(err))
	}
}

func ToJsonErr(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}
