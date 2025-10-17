package handlers

import (
	"encoding/json"
	"libapp/internal/models"
	"libapp/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	//decode json
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
	}

	//call service layer
	createdBook, err := h.service.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//setup response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//send reponse
	json.NewEncoder(w).Encode(createdBook)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	//decode json
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//service layer
	booksList, err := h.service.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//send response
	json.NewEncoder(w).Encode(booksList)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	//extract id from path
	vars := mux.Vars(r)
	id := vars["id"]

	//decode json
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//service layer
	bookDetails, err := h.service.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	//send response
	json.NewEncoder(w).Encode(bookDetails)
}
