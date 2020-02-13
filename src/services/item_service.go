package services

import (
	"github.com/prosline/pl_items_api/src/domain/items"
	"github.com/prosline/pl_util/utils/rest_errors"
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
	Put(items.Item) (*items.Item, *rest_errors.RestErr)
	Delete(items.Item) (*items.Item, *rest_errors.RestErr)
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
func (s *itemService) Get(i string) (*items.Item, *rest_errors.RestErr){
	return nil, nil
}
func (s *itemService) Put(i items.Item) (*items.Item, *rest_errors.RestErr){
	return nil, nil
}
func (s *itemService) Delete(i items.Item) (*items.Item, *rest_errors.RestErr){
	return nil, nil
}

