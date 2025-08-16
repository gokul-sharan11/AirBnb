package main

import (
	"auth_service/app"
	"auth_service/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig(":8080")
	application := app.NewApplication(cfg)
	application.Run()
}