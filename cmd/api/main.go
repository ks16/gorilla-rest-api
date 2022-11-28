package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ks16/gorilla-rest-api/internal/config"
	"github.com/ks16/gorilla-rest-api/internal/router"
	"github.com/ks16/gorilla-rest-api/internal/server"
	"github.com/ks16/gorilla-rest-api/pkg/log"
)

func main() {
	cfg := config.NewConfig()
	logger := log.NewLogger(cfg.Name, cfg.Debug)
	logger.Info().Msgf("starting server on port: %d", cfg.WebServer.Port)

	r := router.NewRouter(logger)
	srv := server.NewServer(cfg, r)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal().Err(err).Msgf("server error")
		}
	}()

	// Graceful web server shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until receive shutdown signal
	<-c
	logger.Info().Msg("received interrupt signal")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	logger.Info().Msg("shutting down web server")
	srv.Shutdown(ctx)

	logger.Info().Msg("shutting down application")
	os.Exit(0)
}
