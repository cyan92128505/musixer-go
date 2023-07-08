package app

import (
	_ "musixer/api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Musixer
// @version 1.0
// @description Musixer API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email cyan92128505@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name access_token
func NewAPP() (app *fiber.App) {
	app = fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	app.Get("/docs/*", swagger.HandlerDefault)

	api := NewAPI()
	app.Mount("/api", api)

	return app
}
