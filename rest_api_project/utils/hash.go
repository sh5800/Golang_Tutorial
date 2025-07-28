package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(passwordBytes),err
}