package app

import (
	"log"

	_ "musixer/api/docs"
	"musixer/api/internal/app/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

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
func NewAPP() (app *fiber.App) {
	app = fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	api := NewAPI()
	app.Mount("/api", api)

	return app
}
