package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"service-api/models"
	"service-api/utils/types"
)

var books []models.Book

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  true,
			Message: "List Books returned successfully",
			Data:    books,
		},
	)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}
	var validate = validator.New()
	if err := validate.Struct(book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  true,
			Message: "Book created successfully",
			Data:    book,
		},
	)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, item := range books {
		if item.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				types.Response{
					Status:  true,
					Message: "Book returned successfully",
					Data:    item,
				},
			)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  false,
			Message: "Book not found with id" + id,
			Data:    nil,
		},
	)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
			var book models.Book
			if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(
					types.Response{
						Status:  false,
						Message: err.Error(),
						Data:    nil,
					},
				)
				return
			}
			book.ID = id
			books = append(books, book)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				types.Response{
					Status:  true,
					Message: "Book updated successfully",
					Data:    book,
				},
			)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  false,
			Message: "Book not found with id" + id,
			Data:    nil,
		},
	)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				types.Response{
					Status:  true,
					Message: "Book deleted successfully",
					Data:    item,
				},
			)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  false,
			Message: "Book not found with id" + id,
			Data:    nil,
		},
	)
}
