package router

import (
	"auth_service/controllers"
	"auth_service/middlewares"
	"auth_service/utils"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middlewares.RequestValidator)
	router.Use(middlewares.RateLimiterMiddleware)
	router.Get("/ping", controllers.PingHandler)
	router.HandleFunc("/api/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.com", "/api/fakestoreservice"))
	UserRouter.Register(router)
	return router
}