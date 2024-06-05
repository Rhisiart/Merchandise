package api

import (
	"fmt"
	"net/http"

	tables "github.com/Rhisiart/Merchandise/internal/db/tables"
	"github.com/Rhisiart/Merchandise/types"
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
	id, err := GetIdFromURL(r, "customerId")

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
	}

	customer := &tables.Customer{
		CustomerId: id,
	}

	queryErr := s.database.Read(r.Context(), customer)

	if queryErr != nil {
		render.Render(w, r, NewError(
			queryErr,
			fmt.Sprintf("The customer with id %d not found.", id),
			http.StatusNotFound))
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

func (s *Server) handlePatchCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromURL(r, "customerId")

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
	}

	data := CustomerRequest{}

	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	customer := &tables.Customer{
		CustomerId: id,
	}

	combine(customer, data)

	queryErr := s.database.Update(r.Context(), customer)

	if queryErr != nil {
		render.Render(w, r, ErrInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}

func (s *Server) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromURL(r, "customerId")

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
	}

	customer := &tables.Customer{
		CustomerId: id,
	}

	dbErr := s.database.Delete(r.Context(), customer)

	if dbErr != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
