package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return pass, err
}
