package repo

import (
	"awesomeProject2/messages/database"
	"awesomeProject2/messages/models"
	"time"
)

func SaveMessage(message *models.Message) error {
	message.CreatedAt = time.Now()
	return database.DBconn.Create(message).Error
}

func GetUserMessages(email string) ([]models.Message, error) {
	var messages []models.Message
	err := database.DBconn.Where(
		"from_email = ? OR to_email = ?",
		email, email).Order("created_at asc").Find(&messages).Error
	return messages, err
}

func MarkMessageAsRead(messageID string) error {
	now := time.Now()
	return database.DBconn.Model(&models.Message{}).
		Where("id = ?", messageID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": now,
		}).Error
}
