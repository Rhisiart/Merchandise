package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/api/customer", func(r chi.Router) {
		r.Get("/", s.handleGetAllCustomers)
		r.Post("/", s.handleCreateCustomer)

		r.Route("/{customerId}", func(r chi.Router) {
			r.Get("/", s.handleGetCustomer)
		})
	})
}
