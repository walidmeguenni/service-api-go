package routes

import (
	"service-api/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouters(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	bookHandler := handlers.NewBookHandler(db)

	router.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	router.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", bookHandler.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
	return router
}
