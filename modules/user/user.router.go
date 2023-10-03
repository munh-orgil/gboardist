package user

import (
	"github.com/gofiber/fiber/v2"

	user_handlers "gboardist/modules/user/handlers"
	"gboardist/session"
)

func SetRoutes(app *fiber.App) {
	var userHandler user_handlers.UserHandler
	userApi := app.Group("/user")
	userApi.Get("/find", session.TokenMiddleware, userHandler.Find)
	userApi.Get("/", session.TokenMiddleware, userHandler.List)
	userApi.Get("/profile", session.TokenMiddleware, userHandler.Profile)
	userApi.Put("/", session.TokenMiddleware, userHandler.Update)
	userApi.Delete("/", session.TokenMiddleware, userHandler.Delete)
	userApi.Put("/change/org", session.TokenMiddleware, userHandler.ChangeOrg)
}
