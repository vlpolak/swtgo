package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword : creates password hashcode
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// HashPassword : checks password hashcode
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
