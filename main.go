package main

import (
	"fmt"
	"net/http"
	"quoter/router"
	"quoter/storage"
)

func main() {
	storage := storage.NewQuoteStorage()  // Create storage instance
	router := router.SetupRouter(storage) // Pass storage to handlers via router

	port := ":8080"
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, router)
}
