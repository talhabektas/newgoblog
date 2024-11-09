package router

import (
	"awesomeProject2/messages/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(handler.WebsocketHandler))

	messages := app.Group("/api/messages")
	messages.Get("/", handler.GetMessages)
	messages.Post("/send", handler.SendMessage)
	messages.Put("/:id/read", handler.MarkMessageAsRead)

}
