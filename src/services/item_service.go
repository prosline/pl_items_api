package services

import (
	"github.com/prosline/pl_items_api/src/domain/items"
	"github.com/prosline/pl_items_api/src/domain/queries"
	"github.com/prosline/pl_util/utils/rest_errors"
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemService struct{}

var (
	ItemService itemServiceInterface = &itemService{}
)

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr){
	if err := item.Save(); err != nil {
		return nil,err
	}
	return &item, nil
}
func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr){
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}
func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr){
	data := items.Item{}
	return data.Search(query)
}

