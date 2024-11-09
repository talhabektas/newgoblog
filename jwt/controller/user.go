package controller

import (
	"awesomeProject2/jwt/authentication"
	"awesomeProject2/jwt/db"
	"awesomeProject2/jwt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func Login(c *gin.Context) {
	objectHandle := gin.H{
		"status": "OK",
		"msg":    "Login",
	}
	var Data Data
	if err := c.ShouldBind(&Data); err != nil {
		log.Println("error")
		c.JSON(400, objectHandle)
		return
	}
	var user models.User
	db.DBconn.First(&user, "email=?", Data.Email)
	if user.ID == 0 {
		objectHandle["msg"] = "user not exist."
		c.JSON(400, objectHandle)
		return
	}

	db.DBconn.First(&user, "username=?", Data.Username)
	if user.ID == 0 {
		objectHandle["msg"] = "user not exist."
		c.JSON(400, objectHandle)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Data.Password))
	if err != nil {
		log.Println("geçersiz parola,tekrar deneyiniz")
		objectHandle["msg"] = "parolalar uyuşmuyor"
		c.JSON(401, objectHandle)
		return
	}

	token, err := authentication.CreateToken(user)
	if err != nil {
		objectHandle["msg"] = "token hatası"
		c.JSON(401, objectHandle)
		return
	}
	objectHandle["token"] = token
	objectHandle["user"] = user
	objectHandle["status"] = "OK"
	objectHandle["msg"] = "başarıyla giriş sağlanmıştır"
	c.JSON(200, objectHandle)

}

func Register(c *gin.Context) {
	objectHandle := gin.H{
		"status": "OK",
		"msg":    "Registered",
	}

	var Data Data
	if err := c.ShouldBind(&Data); err != nil {
		log.Println("error in json")
		objectHandle["msg"] = "HATA"
		c.JSON(400, objectHandle)
		return
	}
	var user models.User
	user.Email = Data.Email
	user.Username = Data.Username
	user.Password = authentication.Hashing(Data.Password)

	result := db.DBconn.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)

		objectHandle["msg"] = "user kaydı zaten önceden tanımlanmıştır."
		c.JSON(400, objectHandle)
		return
	}
	objectHandle["msg"] = "user başarıyla kaydedilmiştir."
	c.JSON(201, objectHandle)
}

func LogOut(c *gin.Context) {
	objectHandle := gin.H{
		"status": "OK",
		"msg":    "Logged Out",
	}
	c.JSON(200, objectHandle)
}
func Tokens(c *gin.Context) {
	objectHandle := gin.H{
		"status": "OK",
		"msg":    "tokennnneddd",
	}
	c.JSON(200, objectHandle)
}
