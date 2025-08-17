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

func CheckPasswordHash(plainPassword string, hash string) bool {
	fmt.Println(plainPassword)
	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
	return err == nil
}