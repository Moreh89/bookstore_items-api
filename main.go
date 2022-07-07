package main

import (
	"github.com/Moreh89/bookstore_items-api/controllers"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {
	router.handleFunc("/items", controllers.Create)
}
