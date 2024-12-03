package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"receipt-processor/internal/models"
)

// ErrorMiddleware captures panics and handles errors gracefully.
func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// Log the panic with request details
				log.Printf("Panic occurred: %v | Path: %s | Method: %s | Client: %s",
					rec, r.URL.Path, r.Method, r.RemoteAddr)

				// Respond with a 500 Internal Server Error
				respondWithError(w, http.StatusInternalServerError, "Internal Server Error", nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
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
