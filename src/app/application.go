package app

import (
	"github.com/gorilla/mux"
	"github.com/prosline/pl_items_api/src/clients/elasticsearch"
	"github.com/prosline/pl_logger/logger"
	"net/http"
)

var (
	router = mux.NewRouter()
)
func StartApplication(){
	elasticsearch.Init()
	logger.Info("Starting Application on port 8082......")
	URLMapping()
	http.ListenAndServe(":8082", router)

}
