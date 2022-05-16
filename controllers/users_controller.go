package controllers

import (
	"bookstore/items-api/domain/items"
	"bookstore/items-api/services"
	"fmt"
	"net/http"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request) 
}

type usersController struct {}

func (c* usersController) Create(w http.ResponseWriter, r *http.Request) {
	// authenticate user using oauth microservice

	item := items.Item{
		Seller: 1, // get the seller using oauth oauth.GetCallerId
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}
	fmt.Println(result)
}

func (c* usersController) Get(w http.ResponseWriter, r *http.Request) {

}
