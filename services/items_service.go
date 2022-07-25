package services

import (
	"net/http"

	"github.com/moreh89/bookstore_items-api/domain/items"
	"github.com/moreh89/bookstore_items-api/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, errors.RestError)
	Get(string) (*items.Item, errors.RestError)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, errors.RestError) {
	return nil, errors.RestError{Message: "implement me!", Status: http.StatusNotImplemented, Error: "not_implemented"}
}
func (s *itemsService) Get(string) (*items.Item, errors.RestError) {
	return nil, errors.RestError{Message: "implement me!", Status: http.StatusNotImplemented, Error: "not_implemented"}
}
