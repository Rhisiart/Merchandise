package api

import (
	"net/http"
	"strconv"

	tables "github.com/Rhisiart/Merchandise/internal/db/tables"
	"github.com/Rhisiart/Merchandise/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) handleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []types.Table
	customer := &tables.Customer{}

	err := s.database.ReadAll(r.Context(), customer, &customers)

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewListResponse(customers))
}

func (s *Server) handleGetCustomer(w http.ResponseWriter, r *http.Request) {
	idParameter := chi.URLParam(r, "customerId")
	id, err := strconv.Atoi(idParameter)

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
	}

	customer := &tables.Customer{
		CustomerId: id,
	}

	queryErr := s.database.Read(r.Context(), customer)

	if queryErr != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Render(w, r, NewCustomerResponse(customer))
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

	render.Render(w, r, NewCustomerResponse(customer))
}
