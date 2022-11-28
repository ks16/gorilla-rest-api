package router

import (
	"github.com/gorilla/mux"
	"github.com/ks16/gorilla-rest-api/internal/handlers"
	"github.com/ks16/gorilla-rest-api/pkg/log"
	"github.com/ks16/gorilla-rest-api/pkg/middleware"
)

func NewRouter(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware(logger))

	router.HandleFunc("/_health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")
	router.HandleFunc("/pong", handlers.Pong).Methods("GET")

	return router
}
