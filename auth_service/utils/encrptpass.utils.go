package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error Hashing Password:", err)
		return "", err
	}
	return string(bytes), err
}

func CheckPasswordHash(plainPassord string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassord))
	return err == nil
}