// handlers.go
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

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
