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
// commands.go
package main

import "fmt"

// Simulated functions for demonstration purposes
func dial(value string) error {
	fmt.Printf("Dialing %s\n", value)
	return nil
}

func answer() error {
	fmt.Println("Answering")
	return nil
}

func hangup() error {
	fmt.Println("Hanging up")
	return nil
}

func sendDtmf(value string) error {
	fmt.Printf("Sending DTMF %s\n", value)
	return nil
}
