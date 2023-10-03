package log

import (
	"gboardist/session"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	var logHandler LogHandler
	logApi := app.Group("/log")
	logApi.Get("/find", session.TokenMiddleware, logHandler.Find)
	logApi.Get("/", session.TokenMiddleware, logHandler.List)

	app.Get("/test", session.TokenMiddleware, logHandler.Test)
}
