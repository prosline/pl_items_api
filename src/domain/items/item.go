package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prosline/pl_items_api/src/clients/elasticsearch"
	"github.com/prosline/pl_items_api/src/domain/queries"
	"github.com/prosline/pl_util/utils/rest_errors"
	"net/http"
	"strings"
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
	PlainText string `json:"plain_text"`
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
func (i *Item) Save() rest_errors.RestErr{
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
			rest_errors.NewRestError("Implement me", http.StatusNotImplemented, "not_implemented", nil)
		}
	i.Id = result.Id
	return nil
}
func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems,typeItem,i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("No item found with Id = %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("Error when trying to get the item by Id", i.Id), errors.New("Database Error"))
	}
	if !result.Found{
		return rest_errors.NewNotFoundError(fmt.Sprintf("No items found with id = %s", i.Id))
	}
	bytes, err := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes,i); err != nil {
		return rest_errors.NewInternalServerError("Error parsing JSON data from Database", errors.New("JSON Error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
