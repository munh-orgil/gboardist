package param

import (
	param_handlers "gboardist/modules/param/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	var paramHandler param_handlers.ParamHandler

	paramApi := app.Group("param")
	paramApi.Get("", paramHandler.List)
	paramApi.Post("", paramHandler.Create)
	paramApi.Put("", paramHandler.Update)
	paramApi.Delete("", paramHandler.Delete)
}
