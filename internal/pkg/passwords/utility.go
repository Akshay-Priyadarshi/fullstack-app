package passwords

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", NewPasswordError(fmt.Errorf("hashing failed: %w", err))
	}
	return string(passwordHash), nil
}

func ComparePassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return NewPasswordError(fmt.Errorf("passwords do not match: %w", err))
	}
	return nil
}
