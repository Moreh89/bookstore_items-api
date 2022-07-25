package app

import (
	"net/http"

	"github.com/moreh89/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
}