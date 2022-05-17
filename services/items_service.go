package services

import (
	"bookstore/items-api/domain/items"
	"bookstore/items-api/utils/errors"
	"time"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(int64) (*items.Item, *errors.RestErr)
	Update(items.Item, int64) (*items.Item, *errors.RestErr)
	Delete(items.Item, int64) *errors.RestErr
	List() ([]*items.Item, *errors.RestErr)
	GetByUserId(int64) ([]*items.Item, *errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) GetByUserId(userId int64) ([]*items.Item, *errors.RestErr) {

	temp := items.Item{
		Seller: userId,
	}

	return temp.GetByUserId()

}

func (s *itemsService) List() ([]*items.Item, *errors.RestErr) {

	temp := items.Item{}

	return temp.List()

}
func (s *itemsService) Delete(item items.Item, clientId int64) *errors.RestErr {

	if item.Id < 1 {
		return errors.NewBadRequestError("itemId cannot be less than 1")
	}

	if err := item.GetById(); err != nil {
		return err
	}

	if clientId != item.Seller {
		return errors.NewBadRequestError("only owner can delete an item")
	}

	if err := item.Delete(); err != nil {
		return err
	}

	return nil

}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {

	if err := item.Validate(); err != nil {
		return nil, err
	}
	item.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")

	if err := item.Create(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(itemId int64) (*items.Item, *errors.RestErr) {
	// return nil, errors.NewInternalServerError("implement me")

	item := items.Item{
		Id: itemId,
	}

	if err := item.GetById(); err != nil {
		return nil, err
	}

	return &item, nil


}
func (s *itemsService) Update(item items.Item, clientId int64) (*items.Item, *errors.RestErr) {
	// return nil, errors.NewInternalServerError("implement me")


	currentItem := items.Item {
		Id: item.Id,
	}

	if err := currentItem.GetById(); err != nil {
		return nil, err
	}

	if clientId != currentItem.Seller {
		return nil, errors.NewBadRequestError("only owner can make updates to an item")
	}

	if len(item.Title) != 0 {
		currentItem.Title = item.Title
	}
	if item.Price != 0 {
		currentItem.Price = item.Price
	}
	if item.Seller != 0 {
		currentItem.Seller = item.Seller
	}
	if item.Stock != 0 {
		currentItem.Stock = item.Stock
	}
	if item.SoldQuantity != 0 {
		currentItem.SoldQuantity = item.SoldQuantity
	}

	if err := currentItem.Update(); err != nil {
		return nil, err
	}


	return &currentItem, nil


}
