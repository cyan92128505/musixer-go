package app

import (
	"fmt"
	"musixer/api/internal/app/controllers"
	middleware "musixer/api/internal/app/middlewares"
	"musixer/api/internal/app/utils"

	"github.com/gofiber/fiber/v2"
)

func NewAPI() (app *fiber.App) {
	app = fiber.New()

	app.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controllers.SignUpUser)
		router.Post("/login", controllers.SignInUser)
		router.Get("/refresh", controllers.RefreshAccessToken)
		router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
	})

	app.Get("/users/me", middleware.DeserializeUser, controllers.GetMe)

	app.Get("/healthcheck", utils.HealthCheck)

	app.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	return app
}
