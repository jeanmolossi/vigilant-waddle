package cmd

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Cors() echo.MiddlewareFunc {
	AllowOrigins := []string{
		"https://jeanmolossi.com.br",
		"https://*.jeanmolossi.com.br",
		"https://lms.jeanmolossi.com.br",
	}
	AllowHeaders := []string{
		echo.HeaderOrigin,
		echo.HeaderContentType,
		echo.HeaderAccept,
		"Authorization",
	}

	if os.Getenv("ENVIRONMENT") == "development" {
		AllowOrigins = append(AllowOrigins, "http://localhost:3000")
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     AllowOrigins,
		AllowHeaders:     AllowHeaders,
		AllowCredentials: true,
	})
}
