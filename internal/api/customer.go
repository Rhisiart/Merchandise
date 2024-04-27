package api

import (
	"net/http"

	tables "github.com/Rhisiart/Merchandise/internal/db/tables"
	"github.com/go-chi/render"
)

type CustomerRequest struct {
	*tables.Customer
}
type CustomerResponse struct {
	*tables.Customer
}

func (crt *CustomerRequest) Bind(r *http.Request) error {
	return nil
}

func (crt *CustomerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleListCustomers(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
}

func (s *Server) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	data := &CustomerRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	customer := &tables.Customer{
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}

	err := s.database.Create(r.Context(), customer)

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Render(w, r, &CustomerResponse{
		Customer: customer,
	})
}
