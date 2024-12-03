package routes

import (
	"net/http"

	"receipt-processor/internal/controllers"
	"receipt-processor/internal/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// NewRouter initializes the router with routes and middleware.
func NewRouter(controller *controllers.ReceiptController, validate *validator.Validate) *mux.Router {
	router := mux.NewRouter()

	// Apply global middleware
	router.Use(middleware.ErrorMiddleware) // Error handling middleware

	// Define routes
	router.HandleFunc("/receipts/process", controller.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", controller.GetPoints).Methods("GET")

	// Status check endpoint
	router.HandleFunc("/status-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Status Check: OK"))
	}).Methods("GET")

	return router
}
