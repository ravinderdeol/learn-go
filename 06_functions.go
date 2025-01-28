// Functions in Go:
// - Functions are blocks of code that execute when called
// - The 'main()' function is special and runs automatically
// - Functions reduce code repetition and solve similar problems with one solution
// - Functions can return one or more values

// Parameters and arguments in Go:
// - Parameters are input variables defined in a function
// - Arguments are the actual values passed to parameters when calling a function
// - The number of arguments must match the number of parameters
// - Go throws an error if the required number of arguments is not provided

// Return values in Go:
// - Functions use 'return' to send values back to the caller
// - The 'return' statement stops the function and sends the specified value
// - Functions specify the type of value they return
// - Returned values can be stored in variables for later use

// Scope in Go:
// - Scope defines where variables and functions are accessible
// - Variables created in a function are limited to that function's scope
// - Global scope allows sharing functions like 'main()' with others
// - Functions cannot directly access variables from other functions
// - To share data between functions use return values and pass them as needed

// Defer in Go:
// - The 'defer' keyword delays a function call until the end of the current scope
// - Use 'defer' for cleanup tasks like closing files or logging
// - 'defer' ensures the function executes even if the code exits early

package main

import "fmt"

const lifespanInWeeks int = 90 * 52

func calculateWeeksLived(age int) int {
	return age * 52
}

func calculateWeeksLeft(weeksLived int) int {
	return lifespanInWeeks - weeksLived
}

func main() {
	var age int
	fmt.Print("Enter your current age: ")
	fmt.Scan(&age)

	weeksLived := calculateWeeksLived(age)
	weeksLeft := calculateWeeksLeft(weeksLived)

	fmt.Println("Weeks lived:", weeksLived)
	fmt.Println("Weeks remaining:", weeksLeft)
}
