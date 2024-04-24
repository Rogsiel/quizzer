package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("Failed to hash Password: %w", err)
    }
    return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
	return fmt.Errorf("Password is incorrect")
    }
    return nil
}
