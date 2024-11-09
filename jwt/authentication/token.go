package authentication

import (
	"awesomeProject2/jwt/models"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type Claims struct {
	Email    string
	UserId   uint
	Username string

	jwt.RegisteredClaims
}

var hidden string = "hidden"

func CreateToken(user models.User) (string, error) {
	claims := Claims{
		user.Email,
		user.ID,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 5)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(hidden))

	if err != nil {
		log.Println("token oluştururken bir hata meydana geldi", err)
		return "", err
	}
	return t, nil
}

func TokenValidation(clientToken string) (claims *Claims, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(hidden), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		msg = err.Error()
		return
	}
	return claims, msg
}
