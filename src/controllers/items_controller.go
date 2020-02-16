package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prosline/pl_items_api/src/domain/items"
	"github.com/prosline/pl_items_api/src/domain/queries"
	"github.com/prosline/pl_items_api/src/services"
	"github.com/prosline/pl_items_api/src/utils/http_utils"
	"github.com/prosline/pl_oauth/oauth"
	"github.com/prosline/pl_util/utils/rest_errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)

}
type itemsController struct{}

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status())
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid access token")
		http_utils.RespondJsonError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		respErr := rest_errors.NewBadRequestError("Invalid Request Body")
		http_utils.RespondJsonError(w,respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody,&itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid Request JSON Body")
		http_utils.RespondJsonError(w, respErr)
		return
	}
	itemRequest.Seller = sellerId

	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondJsonError(w, createErr)
		return
	}
	http_utils.RespondJson(w,http.StatusCreated,result)
}
func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := strings.TrimSpace(vars["id"])
	item, err := services.ItemService.Get(id)
	if err != nil {
		http_utils.RespondJsonError(w,err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}
func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondJsonError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondJsonError(w, apiErr)
		return
	}

	items, searchErr := services.ItemService.Search(query)
	if searchErr != nil {
		http_utils.RespondJsonError(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}
