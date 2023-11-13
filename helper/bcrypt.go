package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	salt := 10
	plain_password := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(plain_password, salt)

	return string(hash)
}

func ComparePassword(h, password []byte) bool {
	hash, pass := []byte(h), []byte(password)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
