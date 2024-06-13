package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetIdFromURL(r *http.Request, paramName string) (int, error) {
	id := chi.URLParam(r, paramName)
	return strconv.Atoi(id)
}
