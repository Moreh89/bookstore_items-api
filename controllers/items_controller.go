package controllers

import (
	"net/http"

	"github.com/Moreh89/bookstore_items-api/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest; err != nil {
		// TODO: return error to the caller.
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),

	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}
	fmt.Println(result)
	
}

func Get() (w http.ResponseWriter, r *http.Request) {
	return
}
