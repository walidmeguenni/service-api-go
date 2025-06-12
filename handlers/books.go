package handlers

import (
	"encoding/json"
	"net/http"
	"service-api/models"
	"service-api/utils/types"
	"service-api/validation"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var books []models.Book

type BookHandler struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		DB:       db,
		Validate: validator.New(),
	}
}

func (bh *BookHandler) GetBooks(w http.ResponseWriter, _ *http.Request) {
	var books []models.Book;
	if err:= bh.DB.Find(&books).Error;
	err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.Response{
			Status:  false,
			Message: "Error getting list of books",
			Data:    nil,
		})
		return
	}
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

func (bh *BookHandler)  CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var body validation.CreateBook

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Invalid request body: " + err.Error(),
				Data:    nil,
			},
		)
		return
	}
	var validate = validator.New()
	if err := validate.Struct(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Response{
			Status:  false,
			Message: "Invalid request body: " + err.Error(),
			Data:    nil,
		})
		return
	}

	book := models.Book{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	}

	if err := bh.DB.Create(&book).Error;
	err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.Response{
			Status:  false,
			Message: "Failed to save book to database: " + err.Error(),
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  true,
			Message: "Book created successfully",
			Data:    book,
		},
	)
}

func (bh *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id_64, err := strconv.ParseUint(params["id"], 0, 64);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Invalid book id",
				Data:    nil,
			},
		)
		return
	}
    id := uint(id_64)
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
			Message: "Book not found with id" + params["id"],
			Data:    nil,
		},
	)
}

func (bh *BookHandler)  UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var exitingBook models.Book
	if err := bh.DB.First(&exitingBook, "id = ?", id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Book not found with id" + id,
				Data:    nil,
			},
		)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Invalid request body: " + err.Error(),
				Data:    nil,
			},
		)
		return
	}
	if err := bh.Validate.Struct(updatedBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Validation failed:  " + err.Error(),
				Data:    nil,
			},
		)
		return
	}

	if err:= bh.DB.Model(&exitingBook).Updates(updatedBook).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Failed to update book: " + err.Error(),
				Data:    nil,
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  true,
			Message: "Book updated successfully",
			Data:    exitingBook,
		},
	)
}

func (bh *BookHandler)  DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
    
	var book models.Book
	if err := bh.DB.First(&book, "id = ?", id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Book not found with id" + id,
				Data:    nil,
			},
		)
		return
	}

	if err := bh.DB.Delete(&book).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			types.Response{
				Status:  false,
				Message: "Failed to delete book: " + err.Error(),
				Data:    nil,
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		types.Response{
			Status:  true,
			Message: "Book deleted successfully",
			Data:    book,
		},
	)
	
}
