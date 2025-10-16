package main

import (
	"libapp/internal/api"
	"libapp/internal/api/handlers"
	"libapp/internal/repository/memory"
	"libapp/internal/service"
	"log"
	"net/http"
)

func main() {
	//initialising dependencies
	bookRepo := memory.NewMemoryBookRepo()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	//setup router
	r := api.SetupRoutes(bookHandler)

	//start server
	log.Println("Server starting to serve at port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server Failed %v", err)
	}

}
