// responses.go
package main

import (
	"encoding/json"
	"net/http"
)

// APIResponse represents a simple API response struct
type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// APIError represents an API error response struct
type APIError struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// httpError returns an API error response in JSON format
func httpError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, _ := json.Marshal(APIError{Status: status, Error: http.StatusText(status), Message: message})
	w.Write(jsonData)
}
