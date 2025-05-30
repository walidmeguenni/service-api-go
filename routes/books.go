package routes

import (
	"github.com/gorilla/mux"
	"service-api/handlers"
)

func SetupRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	return router
}
