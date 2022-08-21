package cmd

import (
	"net/http"

	"github.com/jeanmolossi/vigilant-waddle/src/cmd/httputil"
	ah "github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/handler"
	sh "github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/handler"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = validator.NewCustomValidator()

	// Middlewares
	e.Use(Cors())
	e.Use(middleware.RequestID())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/ping", Ping)

	// auth
	e.POST("/auth/login", ah.Login())

	e.GET("/students", sh.GetStudents())
	e.POST("/student", sh.RegisterStudent())
	e.GET("/me", sh.GetMe())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Ping is a health check endpoint.
//
// @summary Ping the server.
// @description A simple health check.
// @tags healthcheck
// @accept json
// @produce json
// @success 200 {object} httputil.PingOk
// @failure 500 {object} httputil.PingInternalServerErr
// @failure 502 {object} httputil.PingInternalServerErr
// @failure 503 {object} httputil.PingInternalServerErr
// @router /ping [get]
func Ping(c echo.Context) error {
	if c.Request().Method != "GET" {
		return c.JSON(http.StatusNotAcceptable, nil)
	}

	return c.JSON(http.StatusOK, httputil.PingOk{Message: "pong"})
}
