package authentication

import "golang.org/x/crypto/bcrypt"

func Hashing(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("error in hash")
	}
	return string(hash)
}
