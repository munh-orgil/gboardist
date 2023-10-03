package main

import (
	"gboardist/modules/auth"
	"gboardist/modules/log"
	"gboardist/modules/param"
	"gboardist/modules/user"

	"github.com/craftzbay/go_grc/v2/helpers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Get("test", func(c *fiber.Ctx) error {
		return helpers.Response(c, "hello")
	})
	auth.SetRoutes(app)
	log.SetRoutes(app)
	param.SetRoutes(app)
	user.SetRoutes(app)
}
