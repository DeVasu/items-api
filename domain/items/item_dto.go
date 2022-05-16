package items

import (
	"bookstore/items-api/utils/errors"
	"strings"
)

type Item struct {
	Id                int64  `json:"id"`
	Title             string  `json:"title"`
	Seller            int64   `json:"seller"`
	Price             float32 `json:"price"`
	Stock 			  int64    `json:"stock"`
	SoldQuantity      int64     `json:"sold_quantity"`
	CreatedAt	string `json:"createdAt"`
}

func (item *Item) Validate() *errors.RestErr {

	item.Title = strings.Trim(item.Title, " ")
	
	if item.Price == 0 {
		return errors.NewBadRequestError("price cannot be zero")
	}
	if len(item.Title) == 0 {
		return errors.NewBadRequestError("title cannot be empty")
	}

	return nil

}
