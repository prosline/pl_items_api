package app

import (
	"github.com/prosline/pl_items_api/src/controllers"
	"net/http"
)

func URLMapping() {
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	//router.HandleFunc("/items/{id:[0-9]+}", controllers.ItemsController.Create).Methods(http.MethodPost)
}
