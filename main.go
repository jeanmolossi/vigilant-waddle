package main

import (
	"github.com/jeanmolossi/vigilant-waddle/docs"
	"github.com/jeanmolossi/vigilant-waddle/src/cmd"
	_ "github.com/swaggo/files"
)

// @termsOfService  github.com/jeanmolossi/vigilant-waddle/terms/
// @contact.name	Jean Molossi
// @contact.url		https://jeanmolossi.com.br/
// @contact.email	jean.carlosmolossi@gmail.com
// @license.name	Apache 2.0
// @license.url		github.com/jeanmolossi/vigilant-waddle/LICENSE
// @securityDefinitions.apikey access_token
// @in 			  header
// @name 		  Authorization
func main() {
	docs.SwaggerInfo.Title = "Vigilant Waddle API"
	docs.SwaggerInfo.Description = "This is a users and session manager API."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"

	cmd.RunServer()
}
