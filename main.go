package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Application details
const (
	ApplicationName    = "Tsip Rest API"
	ApplicationVersion = "1.0.0"
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

// Handler function for GET /api
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters 'command' and 'value'
	command := r.URL.Query().Get("command")
	value := r.URL.Query().Get("value")

	// Validate query parameters
	if command == "" {
		httpError(w, http.StatusBadRequest, "Missing command")
		return
	}

	// Convert command to lowercase for case-insensitive comparison
	command = strings.ToLower(command)

	// Switch on command and call corresponding function
	switch command {
	case "dial":
		if value == "" {
			httpError(w, http.StatusBadRequest, "Missing value for Dial command")
			return
		}
		if err := dial(value); err != nil {
			httpError(w, http.StatusInternalServerError, err.Error())
			return
		}
	case "answer":
		if err := answer(); err != nil {
			httpError(w, http.StatusInternalServerError, err.Error())
			return
		}
	case "hangup":
		if err := hangup(); err != nil {
			httpError(w, http.StatusInternalServerError, err.Error())
			return
		}
	case "senddtmf":
		if value == "" {
			httpError(w, http.StatusBadRequest, "Missing value for SendDtmf command")
			return
		}
		if err := sendDtmf(value); err != nil {
			httpError(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		httpError(w, http.StatusBadRequest, "Invalid command")
		return
	}

	// Return success response
	jsonData, err := json.Marshal(APIResponse{Status: "success", Message: "Command executed successfully"})
	if err != nil {
		httpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers and write JSON data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

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

// httpError returns an API error response in JSON format
func httpError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, _ := json.Marshal(APIError{Status: status, Error: http.StatusText(status), Message: message})
	w.Write(jsonData)
}

func main() {
	// Display application name and version
	fmt.Printf("%s v%s\n", ApplicationName, ApplicationVersion)

	// Print test API commands
	fmt.Println("Test API Commands:")
	fmt.Println("--------------------")
	fmt.Println("1. Dial: http://localhost:8080/api?command=Dial&value=1234567890")
	fmt.Println("2. Answer: http://localhost:8080/api?command=Answer")
	fmt.Println("3. Hangup: http://localhost:8080/api?command=Hangup")
	fmt.Println("4. Send DTMF: http://localhost:8080/api?command=SendDtmf&value=*123#")

	// Create an HTTP handler
	http.HandleFunc("/api", apiHandler)

	// Start the server on port 8080
	fmt.Println("\nServer listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
