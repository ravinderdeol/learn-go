// Pass-by-value:
// - Go is a pass-by-value language
// - Functions get copies of the values passed as arguments
// - Changes in a function do not affect the original values
// - Use addresses and pointers to modify values outside a function
// - Use pass-by-value for simple data or when changes inside the function do not need to persist

// Addresses:
// - Every variable has an address where its value is stored
// - Use the '&' operator followed by a variable name to get its address (e.g. &x)
// - Addresses are often shown in hexadecimal format (e.g. 0x1234abcd)
// - Use addresses to pass the location of a variable instead of its value
// - Addresses help save memory when working with large data structures

// Pointers:
// - Pointers are variables that store the address of a value
// - Use the '*' operator to declare a pointer (e.g. var ptr *int)
// - The '*' means the pointer stores an address
// - The type after '*' tells Go what kind of value is stored at the address (e.g. int or string)
// - You can declare pointers implicitly like other variables
// - Use pointers to share data between functions without copying it
// - Pointers are useful for modifying data in the original location

package main

import "fmt"

func main() {
	var bitcoinPrice int = 100000
	fmt.Println(bitcoinPrice)

	// Pointer stores the address of bitcoinPrice
	// The & operator gets the address of bitcoinPrice
	// Alternative: 'bitcoinPricePointer := &bitcoinPrice'
	var bitcoinPricePointer *int = &bitcoinPrice
	fmt.Println(bitcoinPricePointer)

	// Dereference the pointer to access the value at the address
	// The * operator gets or sets the value at the address
	*bitcoinPricePointer = 90000
	fmt.Println(bitcoinPrice)
}
