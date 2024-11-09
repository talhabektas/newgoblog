package main

import (
	"log"
	"os"

	"awesomeProject2/jwt/db"
	"awesomeProject2/jwt/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error in loading env file.")
	}

	db.DatabaseConnection()
}

func main() {

	sqlDb, err := db.DBconn.DB()

	if err != nil {
		log.Println("Error in getting db conn.")
	}

	defer sqlDb.Close()

	port := os.Getenv("port")

	if port == "" {
		port = "8061"
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Auth-token", "token", "Content-Type"},
		AllowCredentials: true,
	}))

	route.Routes(router)

	log.Fatal(router.Run(":" + port))
}
