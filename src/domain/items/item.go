package items

import (
	"github.com/prosline/pl_items_api/src/clients/elasticsearch"
	"github.com/prosline/pl_util/utils/rest_errors"
	"net/http"
)

type Item struct {
	Id                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}
type Description struct {
	Text string `json:"text"`
	Html string `json:"html"`
}
type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
const (
	indexItems = "items"
	typeItem = "_doc"
)
//func (i *Item) Save() rest_errors.RestErr {
//	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
//	if err != nil {
//		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
//	}
//	i.Id = result.Id
//	return nil
//}
func (i *Item) Save() rest_errors.RestErr{
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
			rest_errors.NewRestError("Implement me", http.StatusNotImplemented, "not_implemented", nil)
		}
	i.Id = result.Id
	return nil
}

