package router

import (
	"awesomeProject2/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.BlogList)
	app.Get("/:id", controllers.BlogRead)
	app.Post("/", controllers.BlogCreate)
	app.Put("/:id", controllers.BlogUpdate)
	app.Delete("/:id", controllers.BlogDelete)
}
