package app

import (
	"net/http"
	"time"
	"fmt"
	"auth_service/router"
	"auth_service/db/repository"
	"auth_service/controllers"
	"auth_service/services"
)

type Config struct {
	Addr string // PORT 
}

type Application struct {
	Config Config 
	Store db.Storage
}

func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
		Store: *db.NewStorage(),
	}
}

func (application *Application) Run () error {
	db := db.NewUserRepository()
	us := services.NewUserService(db)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(*uc)
	server := &http.Server{
		Addr: application.Config.Addr,
		Handler : router.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	} 
	
	fmt.Println("Starting server on port " + application.Config.Addr)

	return server.ListenAndServe();
}