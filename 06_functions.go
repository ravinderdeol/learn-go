// Functions:
// - Functions contain code that runs when called
// - The main() function is special and runs automatically
// - Functions are used to avoid repeating code
// - Use functions to solve similar problems with one solution
// - Functions can return one or more values

// Parameters and arguments:
// - Functions can take input values called parameters
// - Parameters are variables used inside the function
// - Arguments are the actual values you pass to the parameters
// - The number of arguments must match the number of parameters
// - Go gives an error if you do not pass enough arguments

// Return values:
// - Functions can send values back using return
// - The return statement stops the function and sends back a value
// - Functions can specify the type of value they return
// - The returned value can be saved in a variable to use later

// Scope:
// - Scope is where variables and functions can be accessed
// - Variables created in a function can only be used inside that function
// - Global scope allows functions like main() to use other functions
// - Functions cannot access variables from other functions
// - To pass data between functions return values and use them as needed

// Defer:
// - Defer delays a function call until the end of the current scope
// - Defer is useful for tasks like cleanup or logging
// - Defer ensures the function runs no matter how the code exits

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
