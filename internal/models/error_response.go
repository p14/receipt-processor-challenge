package models

// ErrorResponse defines the structure of error responses.
type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}
