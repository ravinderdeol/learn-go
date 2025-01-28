// Basic error handling in Go:
// - Handle errors immediately to maintain program stability
// - Use the built-in 'error' type to represent and manage errors
// - Use basic error handling for non-critical issues where the program can continue running

// Panic and recover in Go:
// - Use 'panic' for critical errors that cannot be recovered
// - Use 'recover' in a deferred function to regain control after a 'panic'
// - Combine 'panic' and 'recover' to handle unexpected issues without crashing
// - Defer ensures cleanup code runs during a 'panic'
// - Use for unrecoverable errors or when the program cannot safely continue without intervention

// Custom errors in Go:
// - Create structs implementing the 'error' interface for custom errors
// - Use custom errors to include details like error codes or descriptions
// - Pass custom errors by reference to maintain state and avoid copying
// - Use when more descriptive and specific error information is needed

// Error wrapping and unwrapping in Go:
// - Use 'fmt.Errorf' to add context to errors for easier debugging
// - Use 'errors.Is' to check if an error matches a specific target error
// - Use 'errors.As' to extract and handle specific error types
// - Wrapping errors preserves the original error for traceability
// - Use for adding context or handling specific error types in complex applications

// Logging and reporting errors in Go:
// - Log errors with timestamps, messages, and stack traces for debugging
// - Use 'fmt.Print' for simple logs, 'log.Fatal' to log and exit, and 'panic' for critical issues
// - Explicitly check and handle error values to avoid ignoring issues
// - Effective logging improves monitoring and debugging in production systems
// - Use for recording errors and program state to facilitate debugging and monitoring

// Best practices for error handling in Go:
// - Always handle errors returned by functions
// - Use 'panic' and 'recover' only for unexpected and unrecoverable conditions
// - Add meaningful context to errors for easier debugging
// - Use custom errors to provide clear and descriptive error information

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Custom error type for insufficient funds
type InsufficientFundsError struct {
	Amount  float64
	Balance float64
}

// Implement the 'Error' method for 'InsufficientFundsError'
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Cannot withdraw %.2f with balance of %.2f: insufficient funds", e.Amount, e.Balance)
}

// Attempt to withdraw money and return an error if there are insufficient funds
func withdraw(balance, amount float64) error {
	if amount > balance {
		return &InsufficientFundsError{Amount: amount, Balance: balance}
	}
	return nil
}

// Simulate reading a file and wrap errors with additional context
func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file '%s': %w", filename, err)
	}
	return nil
}

// Check if the configuration file exists and handle critical errors using panic and recover
func checkConfig(filename string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		panic(fmt.Sprintf("Configuration file '%s' is missing", filename))
	}
	fmt.Println("Configuration file is present.")
}

func main() {
	// Handle a withdrawal with custom error handling
	balance := 100.0
	amount := 200.0
	err := withdraw(balance, amount)
	if err != nil {
		fmt.Println("Withdrawal error:", err)
	} else {
		fmt.Println("Withdrawal successful.")
	}

	// Read a file and handle errors with additional context
	filename := "nonexistent.txt"
	err = readFile(filename)
	if err != nil {
		fmt.Println("File error:", err)

		// Check for specific error types using 'errors.As'
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Println("Detailed error info:", pathErr)
		}
	}

	// Check the configuration file and handle critical errors with panic and recover
	fmt.Println("Checking configuration file...")
	checkConfig("config.json")

	// Log an error for debugging purposes
	log.Println("This is an example of logging an error.")

	// Simulate a panic and demonstrate recovery
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in main:", r)
		}
	}()
	fmt.Println("Simulating a panic...")
	panic("Unexpected critical issue!")

	// Program continues normally after recovery
	fmt.Println("Program continues normally after handling all errors.")
}
