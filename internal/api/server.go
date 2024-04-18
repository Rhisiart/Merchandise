package api

import (
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

func (server *Server) Start() {

}
