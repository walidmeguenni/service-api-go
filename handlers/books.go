package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"service-api/models"
	"service-api/utils"
)

var books []models.Book

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			utils.Response{
				Status:  false,
				Message: "Invalide request body",
				Data:    nil,
			},
		)
		return
	}
	var validate = validator.New()
	if err := validate.Struct(book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Validation failed!",
			Data:    nil,
		})
		return
	}
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(
		utils.Response{
			Status:  true,
			Message: "Book created successfully",
			Data:    book,
		},
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			utils.Response{
				Status:  false,
				Message: "Error in creating Book",
				Data:    nil,
			},
		)
		return
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, item := range books {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Book not found with id: "+id, http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
