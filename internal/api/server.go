package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rhisiart/Merchandise/internal/config"
	"github.com/Rhisiart/Merchandise/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	config   config.HTTPServer
	database *storage.Database
	router   *chi.Mux
}

func NewServer(config config.HTTPServer, database *storage.Database) *Server {
	return &Server{
		config:   config,
		database: database,
		router:   chi.NewRouter(),
	}
}

func (s *Server) Start(ctx context.Context) {
	s.routes()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.router,
	}

	log.Printf("Server listing on port %d", s.config.Port)

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
