package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"receipt-processor/internal/models"
	"receipt-processor/internal/services"

	"github.com/gorilla/mux"

	"github.com/go-playground/validator/v10"
)

// ReceiptController handles HTTP requests related to receipts.
type ReceiptController struct {
	Service  services.ReceiptService
	Validate *validator.Validate
}

// NewReceiptController creates a new ReceiptController with the given service and validator.
func NewReceiptController(service services.ReceiptService, validate *validator.Validate) *ReceiptController {
	return &ReceiptController{
		Service:  service,
		Validate: validate,
	}
}

// ProcessReceipt handles POST /receipts/process
func (c *ReceiptController) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	// Decode the JSON request body into the Receipt struct
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON payload", nil)
		return
	}

	// Validate the struct
	if err := c.Validate.Struct(receipt); err != nil {
		// Collect validation errors
		errorsMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorsMap[err.Field()] = validationErrorMessage(err)
		}

		respondWithError(w, http.StatusBadRequest, "Validation failed", errorsMap)
		return
	}

	// Proceed with processing the receipt
	id, err := c.Service.ProcessReceipt(receipt)
	if err != nil {
		log.Printf("Error processing receipt: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to process receipt", nil)
		return
	}

	// Respond with the generated ID
	response := models.ProcessResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(response)
}

// GetPoints handles GET /receipts/{id}/points
func (c *ReceiptController) GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, err := c.Service.GetPoints(id)
	if err != nil {
		if errors.Is(err, services.ErrReceiptNotFound) {
			respondWithError(w, http.StatusNotFound, "Receipt not found", nil)
		} else {
			log.Printf("Error getting points for receipt %s: %v", id, err)
			respondWithError(w, http.StatusInternalServerError, "Failed to get points", nil)
		}
		return
	}

	// Respond with the calculated points
	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(response)
}

// validationErrorMessage maps validation tags to user-friendly messages.
func validationErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required."
	case "datetime":
		return "Invalid datetime format."
	case "currency":
		return "Must be a valid currency amount with up to two decimal places."
	case "min":
		return "Must contain at least one item."
	default:
		return "Invalid value."
	}
}

// respondWithError sends a standardized error response.
func respondWithError(w http.ResponseWriter, statusCode int, message string, errors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := models.ErrorResponse{
		Message: message,
		Errors:  errors,
	}
	json.NewEncoder(w).Encode(response)
}
