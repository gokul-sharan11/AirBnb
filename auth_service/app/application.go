package app

import (
	"net/http"
	"time"
	"fmt"
	"auth_service/router"
	repo "auth_service/db/repository"
	"auth_service/controllers"
	"auth_service/services"
	dbConfig "auth_service/config/db"
)

type Config struct {
	Addr string // PORT 
}

type Application struct {
	Config Config 
}

func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (application *Application) Run () error {
	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error connecting to database", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)
	server := &http.Server{
		Addr: application.Config.Addr,
		Handler : router.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	} 
	
	fmt.Println("Starting server on port " + application.Config.Addr)

	return server.ListenAndServe();
}