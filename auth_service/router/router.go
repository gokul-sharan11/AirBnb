package router

import (
	"auth_service/controllers"
	"github.com/go-chi/chi/v5"
	"auth_service/middlewares"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middlewares.RequestValidator)
	router.Use(middlewares.RateLimiterMiddleware)
	router.Get("/ping", controllers.PingHandler)
	UserRouter.Register(router)
	return router
}