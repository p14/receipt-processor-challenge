package main

import (
	"fmt"
	"log"
	"net/http"

	"receipt-processor/internal/controllers"
	"receipt-processor/internal/repository"
	"receipt-processor/internal/routes"
	"receipt-processor/internal/services"
	"receipt-processor/internal/validator"
)

func main() {
	// Initialize services
	receiptService := services.NewReceiptService(repository.NewReceiptRepository())

	// Initialize validator
	v := validator.NewValidator()

	// Initialize controllers with dependencies
	receiptController := controllers.NewReceiptController(receiptService, v.Validate)

	// Initialize router with controllers
	router := routes.NewRouter(receiptController, v.Validate)

	// Start the server
	port := "8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
