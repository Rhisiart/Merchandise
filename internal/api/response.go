package api

import (
	"net/http"

	tables "github.com/Rhisiart/Merchandise/internal/db/tables"
	"github.com/Rhisiart/Merchandise/types"
	"github.com/go-chi/render"
)

type omit *struct{}

type CustomerRequest struct {
	*tables.Customer
}
type CustomerResponse struct {
	*tables.Customer

	CustomerId omit `json:"CustomerId,omitempty"`
}

func (crt *CustomerRequest) Bind(r *http.Request) error {
	return nil
}

func (crt *CustomerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCustomerResponse(customer *tables.Customer) *CustomerResponse {
	return &CustomerResponse{
		Customer: customer,
	}
}

func NewListResponse(rows []types.Table) []render.Renderer {
	list := []render.Renderer{}

	for _, row := range rows {
		switch r := row.(type) {
		case *tables.Customer:
			customer := NewCustomerResponse(r)
			list = append(list, customer)

			//do for all the tables
		}

	}

	return list
}
