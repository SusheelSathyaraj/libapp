package api

import (
	"libapp/internal/api/handlers"
	"libapp/internal/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(bookHandler *handlers.BookHandler) *mux.Router {
	//create router
	router := mux.NewRouter()

	//adding logging middleware
	router.Use(middleware.LoggingMiddleware)
	//adding recovery middleware
	router.Use(middleware.RecoveryMiddleware)

	router.HandleFunc("/books", bookHandler.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", bookHandler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods(http.MethodGet)

	return router
}
