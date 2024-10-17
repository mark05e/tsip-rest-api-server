package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Check if a process named "tsip.exe" is running and get its full path
func isProcessRunning(processName string) (bool, string, error) {
	// Execute tasklist command to find the process (Windows-specific)
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("IMAGENAME eq %s", processName))

	// Debug: Print the full command being executed
	fmt.Println("Executing command:", strings.Join(cmd.Args, " "))

	// Buffers to capture both standard output and error output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command and capture any errors
	err := cmd.Run()
	if err != nil {
		// Debug: Print the standard error output in case of failure
		fmt.Println("Error executing command:", err)
		fmt.Println("Standard error output:", stderr.String())
		return false, "", fmt.Errorf("command failed with error: %v, stderr: %s", err, stderr.String())
	}

	// Debug: Print the standard output
	output := out.String()
	// fmt.Println("Command output:", output)

	// Convert output and process name to lowercase for a case-insensitive comparison
	lowerOutput := strings.ToLower(output)
	lowerProcessName := strings.ToLower(processName)

	// Check if the process name exists in the output
	if strings.Contains(lowerOutput, lowerProcessName) {
		// Get the full path of the process
		pathCmd := exec.Command("wmic", "process", "where", fmt.Sprintf("name='%s'", processName), "get", "ExecutablePath")
		var pathOut bytes.Buffer
		pathCmd.Stdout = &pathOut
		err = pathCmd.Run()
		if err != nil {
			return true, "", err // Process is running, but path retrieval failed
		}

		pathOutput := strings.TrimSpace(pathOut.String())
		// Remove header "ExecutablePath" and clean up the result
		lines := strings.Split(pathOutput, "\n")
		if len(lines) > 1 {
			executablePath := strings.TrimSpace(lines[1])
			return true, executablePath, nil
		}
		return true, "", nil // Process is running, but no path found
	}
	return false, "", nil // Process is not running
}
