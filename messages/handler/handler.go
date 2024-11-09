package handler

import (
	"awesomeProject2/messages/models"
	"awesomeProject2/messages/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"sync"
)

var (
	clients = make(map[string]*websocket.Conn)
	mutex   = sync.RWMutex{}
)

func WebsocketHandler(c *websocket.Conn) {
	email := c.Query("email")
	if email == "" {
		return
	}

	mutex.Lock()
	clients[email] = c
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, email)
		mutex.Unlock()
	}()

	for {
		messageType, _, err := c.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			break
		}
	}
}

func GetMessages(c *fiber.Ctx) error {
	userEmail := c.Query("email")
	if userEmail == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	messages, err := repo.GetUserMessages(userEmail)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch messages",
		})
	}

	return c.JSON(fiber.Map{
		"messages": messages,
	})
}

func SendMessage(c *fiber.Ctx) error {
	message := new(models.Message)

	if err := c.BodyParser(message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid message data",
		})
	}

	if err := repo.SaveMessage(message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save message",
		})
	}

	mutex.RLock()
	if conn, ok := clients[message.ToEmail]; ok {
		conn.WriteJSON(message)
	}
	mutex.RUnlock()

	return c.Status(fiber.StatusCreated).JSON(message)
}

func MarkMessageAsRead(c *fiber.Ctx) error {
	messageID := c.Params("id")

	if err := repo.MarkMessageAsRead(messageID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to mark message as read",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
