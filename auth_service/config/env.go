package config 

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
)

func Load(){
	err :=godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetString (key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func GetInt (key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return i
}

func GetBool (key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return b
}