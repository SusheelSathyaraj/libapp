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

	router.HandleFunc("/books", bookHandler.CreateBook).Methods(http.MethodPost)

	return router
}
