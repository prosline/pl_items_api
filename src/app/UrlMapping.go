package app

import (
	"github.com/prosline/pl_items_api/src/controllers"
	"net/http"
)

func URLMapping() {
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}
