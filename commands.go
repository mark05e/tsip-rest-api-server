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
