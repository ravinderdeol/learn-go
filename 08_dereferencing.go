// Dereferencing in Go:
// - Dereferencing retrieves or modifies the value at a memory address
// - Use the '*' operator on a pointer to access or update the value it points to
// - Example: '*ptr = 10' updates the value at the address stored in 'ptr'
// - Dereferencing is efficient for updating shared data or large structures

// Pointers as function arguments in Go:
// - Use pointers to modify a variable's value in a different scope
// - Pass a pointer to a function to allow changes to the original variable
// - Inside the function, use the '*' operator to dereference and update the value
// - Use the '&' operator to pass a variable's address as a pointer to a function
// - Pointers allow functions to modify the original variable by sharing its address

// Define a function that takes a pointer as an argument
func brainwash(saying *string) {

	// Update the value at the address using '*'
	*saying = "Beep Boop."
}

func main() {

	// Declare a variable with an initial value
	greeting := "Hello there!"

	// Pass the address of the variable to the function using '&'
	brainwash(&greeting)

	// Print the updated value of the variable
	fmt.Println("greeting is now:", greeting)
}
