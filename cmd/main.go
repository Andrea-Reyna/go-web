package main

import (
	"log"
	"os"

	"github.com/Andrea-Reyna/go-web/cmd/docs"
	"github.com/Andrea-Reyna/go-web/cmd/server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost/8080
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	server := gin.New()

	router := handlers.Router{
		Engine: server,
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Setup()

	router.Engine.Run()
}
