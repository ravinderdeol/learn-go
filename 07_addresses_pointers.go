// Pass-by-value in Go:
// - Go uses pass-by-value for function arguments
// - Functions receive copies of the values passed as arguments
// - Changes made inside a function do not affect the original values
// - Use addresses and pointers to modify values outside the function
// - Pass-by-value is suitable for simple data or when changes are not needed outside the function

// Addresses in Go:
// - Every variable has a memory address where its value is stored
// - Use the '&' operator to get the address of a variable (e.g. '&x')
// - Addresses are shown in hexadecimal format (e.g. '0x1234abcd')
// - Passing addresses lets you share the location of a variable instead of its value
// - Using addresses saves memory when working with large data structures

// Pointers in Go:
// - Pointers are variables that store the address of another value
// - Use the '*' operator to declare a pointer (e.g. 'var ptr *int')
// - The '*' indicates the variable is a pointer storing an address
// - The type after '*' specifies the kind of value stored at the address (e.g. 'int' or 'string')
// - Pointers can be declared implicitly like other variables
// - Use pointers to share data between functions without copying it
// - Pointers allow you to modify data at its original memory location

package main

import "fmt"

func main() {
	var bitcoinPrice int = 100000
	fmt.Println(bitcoinPrice)

	// Declare a pointer to store the address of 'bitcoinPrice' using '&'
	var bitcoinPricePointer *int = &bitcoinPrice
	fmt.Println(bitcoinPricePointer)

	// Dereference the pointer with '*' to update the value at the address
	*bitcoinPricePointer = 90000
	fmt.Println(bitcoinPrice)
}
