package handler

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/producer/adapter"
	"github.com/jeanmolossi/vigilant-waddle/src/core/modules/producer/factory"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/http_error"

	"github.com/labstack/echo/v4"
)

// RegisterProducer godoc
//
// @Summary Register a producer
// @Description Register a producer
// @ID register-producer
// @Tags producer
// @Produce json
// @Param producer body adapter.RegisterProducer true "Producer"
// @Success 201 {object} HttpNewProducer
// @Failure 400 {object} http_error.HTTPBadRequestError "Bad request"
// @Failure 409 {object} http_error.HTTPError "User with that email already exists"
// @Failure 500 {object} http_error.HTTPError "An error occurred"
// @Router /producer [post]
func RegisterProducer() echo.HandlerFunc {
	db := database.GetConnection()
	usecase := factory.RegisterProducer(db)

	return func(c echo.Context) error {
		producerInput := new(adapter.RegisterProducer)

		if err := c.Bind(producerInput); err != nil {
			return http_error.Handle(c, err)
		}

		if err := c.Validate(producerInput); err != nil {
			return http_error.Handle(c, err)
		}

		s := producer.NewProducer(
			producer.WithEmail(producerInput.Email),
			producer.WithPassword(producerInput.Password),
		)

		producer, err := usecase(s)

		if err != nil {
			return http_error.Handle(c, err)
		}

		return c.JSON(http.StatusCreated, NewHttpNewStudent(producer))
	}
}

// Http responses

// HttpProducer is a producer representation for http response
type HttpProducer struct {
	ID    string `json:"id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Email string `json:"email,omitempty" example:"john@doe.com" format:"email"`
	Scope string `json:"scopes,omitempty" example:"producer"`
}

// HttpNewProducer is a producer representation for http response
type HttpNewProducer struct {
	Data HttpProducer `json:"data"`
}

// NewHttpNewStudent creates a new HttpNewProducer
func NewHttpNewStudent(s producer.Producer) *HttpNewProducer {
	return &HttpNewProducer{
		Data: HttpProducer{
			s.GetID(),
			s.GetEmail(),
			string(s.GetScope()),
		},
	}
}
