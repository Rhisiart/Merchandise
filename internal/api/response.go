package api

import (
	"net/http"

	tables "github.com/Rhisiart/Merchandise/internal/storage/tables"
	"github.com/Rhisiart/Merchandise/types"
	"github.com/go-chi/render"
)

//type omit *struct{}

type CustomerRequest struct {
	*tables.Customer
}
type CustomerResponse struct {
	*tables.Customer

	StatusCode int `json:"-"`
	//CustomerId omit `json:"CustomerId,omitempty"`
}

func (cr *CustomerRequest) Bind(r *http.Request) error {
	return nil
}

func (cr *CustomerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	//render.Status(r, cr.StatusCode)
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
