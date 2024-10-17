// main.go
package main

import (
	"fmt"
	"net/http"
)

// Application details
const (
	ApplicationName    = "Tsip Rest API"
	ApplicationVersion = "1.0.0"
)

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
