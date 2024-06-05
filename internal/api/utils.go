package api

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetIdFromURL(r *http.Request, paramName string) (int, error) {
	id := chi.URLParam(r, paramName)
	return strconv.Atoi(id)
}

func combine(dest interface{}, src interface{}) {
	destValue := reflect.ValueOf(dest).Elem()
	srcValue := reflect.ValueOf(src)

	for i := 0; i < destValue.NumField(); i++ {
		destField := destValue.Field(i)
		srcField := srcValue.FieldByName(destValue.Type().Field(i).Name)

		if srcField.IsValid() && !srcField.IsZero() {
			destField.Set(srcField)
		}
	}
}
