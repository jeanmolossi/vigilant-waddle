package http_error

import (
	"errors"
	"net/http"

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
	case errors.As(err, &validationErr):
		return c.JSON(http.StatusBadRequest, err)
	case errors.As(err, &bindErr):
		return c.JSON(http.StatusBadRequest, err)
	default:
		return c.JSON(http.StatusInternalServerError, ToJsonErr(err))
	}
}

func ToJsonErr(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}
