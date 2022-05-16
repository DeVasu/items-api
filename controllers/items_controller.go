package controllers

import (
	"bookstore/items-api/domain/items"
	"bookstore/items-api/services"
	"bookstore/items-api/utils"
	"bookstore/items-api/utils/errors"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request) 
	Update(w http.ResponseWriter, r *http.Request) 
	Delete(w http.ResponseWriter, r *http.Request) 
	List(w http.ResponseWriter, r *http.Request) 
}

type itemsController struct {}

func (c* itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	// authenticate user using oauth microservice
	enableCors(&w)
	itemId := mux.Vars(r)["itemId"]
	itemIdInt, intErr := strconv.Atoi(itemId)
	if intErr != nil {
		tempErr := errors.NewBadRequestError("itemId could not be parsed")
		utils.RespondError(w, *tempErr)
		return
	}
	userId, err := utils.AuthenticateToken(r)
	if err != nil {
		utils.RespondError(w, *err)
		return
	}
	item := items.Item{
		Id: int64(itemIdInt),
	}
	err = services.ItemsService.Delete(item, userId)
	if err != nil {
		utils.RespondError(w, *err)
		return
	}
	x := make(map[string]interface{})
	x["message"] = "successfully deleted"
	utils.RespondJson(w, http.StatusCreated, x)
}

func (c* itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// authenticate user using oauth microservice
	enableCors(&w)


	userId, err := utils.AuthenticateToken(r)
	if err != nil {
		utils.RespondError(w, *err)
		return
	}


	item := items.Item{}

	if err := json.NewDecoder(r.Body).Decode(&item) ; err != nil {
		badErr := errors.NewBadRequestError("something wrong with the input")
		utils.RespondError(w, *badErr)
		return
	}

	item.Seller = userId

	result, err := services.ItemsService.Create(item)
	if err != nil {
		utils.RespondError(w, *err)
		return
	}
	utils.RespondJson(w, http.StatusCreated, result)
}
func (c* itemsController) Update(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// authenticate user using oauth microservice
	itemId := mux.Vars(r)["itemId"]
	itemIdInt, intErr := strconv.Atoi(itemId)
	if intErr != nil {
		tempErr := errors.NewBadRequestError("itemId could not be parsed")
		utils.RespondError(w, *tempErr)
		return
	}

	userId, err := utils.AuthenticateToken(r)
	if err != nil {
		utils.RespondError(w, *err)
		return
	}

	item := items.Item{}

	if err := json.NewDecoder(r.Body).Decode(&item) ; err != nil {
		badErr := errors.NewBadRequestError("something wrong with the input" + err.Error())
		utils.RespondError(w, *badErr)
		return
	}

	item.Id = int64(itemIdInt)

	result, err := services.ItemsService.Update(item, userId)
	if err != nil {
		utils.RespondError(w, *err)
		return
	}
	utils.RespondJson(w, http.StatusCreated, result)
}

func (c* itemsController) Get(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	
	itemId := mux.Vars(r)["itemId"]
	itemIdInt, intErr := strconv.Atoi(itemId)
	if intErr != nil {
		tempErr := errors.NewBadRequestError("itemId could not be parsed")
		utils.RespondError(w, *tempErr)
		return
	}

	item, err := services.ItemsService.Get(int64(itemIdInt))
	if err != nil {
		utils.RespondError(w, *err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.RespondJson(w, http.StatusOK, item)
}
func (c* itemsController) List(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)


	items, err := services.ItemsService.List()
	if err != nil {
		utils.RespondError(w, *err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.RespondJson(w, http.StatusOK, items)
}

func Ping(w http.ResponseWriter, r*http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
