package hasher

import (
	"github.com/vlpolak/swtgo/logger"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword : creates password hashcode
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.ErrorLogger("Creating password failed", err)
	}
	logger.CommonLogger("Password created", password)
	return string(bytes), err
}

// HashPassword : checks password hashcode
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logger.ErrorLogger("Password check failed", err)
	}
	logger.CommonLogger("Password checked", password)
	return err == nil
}
