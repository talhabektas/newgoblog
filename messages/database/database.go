package database

import (
	"awesomeProject2/messages/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DBconn *gorm.DB

func DatabaseConnection() {

	dbconnection := os.Getenv("dbconnection")

	db, err := gorm.Open(mysql.Open(dbconnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("error in db connection")
	}
	log.Println("very good db connection")
	db.AutoMigrate(new(models.Message))
	DBconn = db
}
