package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/cleopatra/pkg/auth"
)

func LimitPhoneNumbersPerRequest(c *fiber.Ctx) error {
	if auth.IsLoggedIn(c) {
		return c.Redirect("/")
	}
	return c.Next()
}