package api

import (
	"net/http"
	"strconv"

	"github.com/Rhisiart/Merchandise/internal/db"
	tables "github.com/Rhisiart/Merchandise/internal/db/tables"
)

type Api struct {
	port     string
	database db.Database
}

func NewApi(port string, database db.Database) *Api {
	return &Api{
		port:     port,
		database: database,
	}
}

func (api *Api) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/customer", handleHttp(api.HandleHttpRequests))
}

func (api *Api) HandleHttpRequests(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		pathValue := r.PathValue("id")
		id, err := strconv.Atoi(pathValue)

		if err != nil {

		}

		c := &tables.Customer{
			CustomerId: id,
		}
		api.database.Read(c)
	case "POST":
		api.database.Create()
	case "PUT":
	case "PATCH":
	case "DELETE":
	}
}
