package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/cleopatra/rest/middlewares"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("api").Use(middlewares.AuthApi())
	web := app.Group("")
	ApiRoutes(api)
	WebRoutes(web)
}
