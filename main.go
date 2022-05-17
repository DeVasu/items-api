package main

import (
	"bookstore/items-api/controllers"
	_ "bookstore/items-api/datasources"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {

	//ping system
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet)


	// get items and create items
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items", controllers.ItemsController.List).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodOptions)

	//get items by userId
	router.HandleFunc("/items/mine", controllers.ItemsController.GetByUserId).Methods(http.MethodGet)
	router.HandleFunc("/items/mine", controllers.ItemsController.GetByUserId).Methods(http.MethodOptions)

	//CRUD for items
	router.HandleFunc("/items/{itemId}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/items/{itemId}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/{itemId}", controllers.ItemsController.Update).Methods(http.MethodPut)



	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8000",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	// ch := make(chan os.Signal)

	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		panic(err)
	// 	}	
	// }

}
