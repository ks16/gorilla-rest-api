package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ks16/gorilla-rest-api/internal/config"
	"net/http"
	"time"
)

func NewServer(cfg *config.Config, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.WebServer.Port),
		WriteTimeout: time.Second * time.Duration(cfg.WebServer.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.WebServer.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.WebServer.IdleTimeout),
		Handler:      router,
	}
}
