package router 

import (
	"github.com/go-chi/chi/v5"
	"auth_service/controllers"
	// "auth_service/middlewares"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middlewares.RequestValidator)
	router.Get("/ping", controllers.PingHandler)
	UserRouter.Register(router)
	return router
}