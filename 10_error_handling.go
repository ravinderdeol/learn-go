// Basic Error Handling:
// - Handle errors immediately to maintain program stability.
// - Use the built-in 'error' type to represent and manage errors.
// - Use: For non-critical issues where the program can continue running after handling the error.

// Panic and Recover:
// - Use 'panic' for critical errors that cannot be recovered.
// - Use 'recover' in a deferred function to regain control after a 'panic'.
// - Combine 'panic' and 'recover' to handle unexpected issues without crashing the program.
// - Defer ensures cleanup code runs during a 'panic'.
// - Use: For unrecoverable errors or situations where the program cannot safely continue without intervention.

// Custom Errors:
// - Create structs that implement the 'error' interface for custom errors.
// - Use custom errors to add details like error codes or descriptions.
// - Pass custom errors by reference to maintain state and avoid copying.
// - Use: When you need more descriptive and specific error information than the default 'error' type provides.

// Error Wrapping and Unwrapping:
// - Use 'fmt.Errorf' to add context to errors for easier debugging.
// - Use 'errors.Is' to check if an error matches a specific target error.
// - Use 'errors.As' to extract and handle specific error types.
// - Wrapping errors preserves the original error for traceability.
// - Use: For adding extra context to errors or working with specific error types in complex applications.

// Logging and Reporting:
// - Log errors with timestamps, messages, and stack traces for debugging.
// - Use 'fmt.Print' for simple logs, 'log.Fatal' to log and exit, and 'panic' for critical issues.
// - Explicitly check and handle error values to prevent ignoring issues.
// - Effective logging improves production monitoring and debugging.
// - Use: For recording errors and program state, especially in production systems, to facilitate debugging and monitoring.

// Best Practices:
// - Always handle errors returned by functions.
// - Use 'panic' and 'recover' only for unexpected, unrecoverable conditions.
// - Add meaningful context to errors to make debugging easier.
// - Use custom errors for clear and descriptive error information.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Custom error type for insufficient funds.
type InsufficientFundsError struct {
	Amount  float64
	Balance float64
}

// Error method implements the error interface for InsufficientFundsError.
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Cannot withdraw %.2f with balance of %.2f: insufficient funds", e.Amount, e.Balance)
}

// withdraw attempts to deduct money from the account and returns an error if insufficient funds.
func withdraw(balance, amount float64) error {
	// Custom Errors: Adding context to an error with a custom type.
	if amount > balance {
		return &InsufficientFundsError{Amount: amount, Balance: balance}
	}
	return nil
}

// readFile simulates reading a file and wraps errors with additional context.
func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// Error Wrapping: Adding more context to the original error.
		return fmt.Errorf("failed to open file '%s': %w", filename, err)
	}
	return nil
}

// checkConfig ensures the configuration file exists, using panic and recover for critical errors.
func checkConfig(filename string) {
	defer func() {
		// Panic and Recover: Catching panic and recovering to continue execution.
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate missing configuration file as a critical error.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Panic: Trigger a panic for a critical, unrecoverable issue.
		panic(fmt.Sprintf("Configuration file '%s' is missing", filename))
	}
	fmt.Println("Configuration file is present.")
}

func main() {
	// Step 1: Handle a withdrawal with custom error handling.
	balance := 100.0
	amount := 200.0
	err := withdraw(balance, amount)
	if err != nil {
		// Basic Error Handling: Checking and responding to errors.
		fmt.Println("Withdrawal error:", err)
	} else {
		fmt.Println("Withdrawal successful.")
	}

	// Step 2: Wrap and unwrap errors while reading a file.
	filename := "nonexistent.txt"
	err = readFile(filename)
	if err != nil {
		fmt.Println("File error:", err)

		// Error Unwrapping: Checking for specific error types.
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Println("Detailed error info:", pathErr)
		}
	}

	// Step 3: Use panic and recover for critical errors.
	fmt.Println("Checking configuration file...")
	checkConfig("config.json")

	// Step 4: Log an error for debugging purposes.
	log.Println("This is an example of logging an error.")

	// Step 5: Simulate a panic and demonstrate recovery.
	defer func() {
		// Panic and Recover: Gracefully handling unexpected issues.
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in main:", r)
		}
	}()
	fmt.Println("Simulating a panic...")
	panic("Unexpected critical issue!")

	// Step 6: Program continues after recovery.
	fmt.Println("Program continues normally after handling all errors.")
}
