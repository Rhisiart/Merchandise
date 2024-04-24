package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rhisiart/Merchandise/internal/config"
	"github.com/Rhisiart/Merchandise/internal/db"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	config   config.HTTPServer
	database *db.Database
	router   *chi.Mux
}

func NewServer(config config.HTTPServer, database *db.Database) *Server {
	return &Server{
		config:   config,
		database: database,
		router:   chi.NewRouter(),
	}
}

func (s *Server) Start() {
	s.routes()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.router,
	}

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		log.Printf("shutdown the server:")
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")
}
