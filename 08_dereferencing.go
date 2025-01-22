// Dereferencing:
// - Dereferencing accesses or changes the value at an address
// - Use the * operator on a pointer to get or set the value it points to
// - Example: *ptr = 10 changes the value at the address stored in ptr
// - Dereferencing is useful for updating shared data or large structures efficiently

// Pointers as function arguments:
// - Use pointers to change a variable's value in a different scope
// - Pass a pointer to a function to modify the original variable
// - In the function, use the * operator to dereference the pointer and update the value
// - To pass a pointer, use the & operator to get the variable's address
// - Pointers allow changes to the original variable because they share the same address

// Define a function that takes a pointer as an argument
func brainwash(saying *string) {

	// Dereference the pointer to update the value at the address
	*saying = "Beep Boop."
}

func main() {

	// Define a variable with an initial value
	greeting := "Hello there!"

	// Pass the variable's address to the function using the & operator
	brainwash(&greeting)

	// Print the updated value of the variable
	fmt.Println("greeting is now:", greeting)
}
