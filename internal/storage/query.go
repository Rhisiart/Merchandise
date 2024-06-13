package storage

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Rhisiart/Merchandise/enum"
)

type Query struct {
	Statement enum.Statement
	Table     string
	Row       interface{}
}

func NewQuery(statement enum.Statement, table string, row interface{}) *Query {
	return &Query{
		Statement: statement,
		Table:     table,
		Row:       row,
	}
}

func (q Query) Expression() string {
	value := reflect.ValueOf(q.Row).Elem()
	t := value.Type()

	var result []string

	for i := 1; i < value.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := value.Field(i).Interface()

		if fieldValue == nil {
			continue
		}

		result = append(result, fmt.Sprintf("%s = $%d", fieldName, i+1))
	}

	return strings.Join(result, ", ")
}

func (q Query) Values() ([]interface{}, error) {
	v := reflect.ValueOf(q.Row)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("the row must be a pointer to a struct")
	}

	v = v.Elem()

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values, nil
}
