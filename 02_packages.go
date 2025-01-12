// Code organisation:
// - Go programs use packages to group related code.
// - A package is defined by the package declaration in a '.go' file.
// - Files with the same package declaration belong to the same package.
// - A folder in Go usually represents a package and holds its files.

// Packages:
// - A package organizes and reuses code in Go.
// - The package declaration tells Go which package a file belongs to.
// - The main package defines the entry point of a Go application.
// - Programs with a main function in the main package are compiled into an executable.
// - The executable starts by running the main function.

// Documentation:
// - Go has a program called go doc to view documentation.
// - Use go doc <package> to see details about a package (e.g. go doc time).
// - Use go doc <package>.<function> to see details about a function (e.g. go doc time.Now).

// Imports:
// - Imports let you use code from other packages.
// - Use parentheses to import multiple packages.
// - Use an alias to rename a package like t for time.

// Functions:
// - The func keyword declares a function.
// - A function name follows the func keyword (e.g. main).
// - Function code is written inside curly braces

package main

import (
	"fmt"
	t "time"
)

func main() {

	// Prints 'Hello, World!' using the Println function from the fmt package.
	fmt.Println("Hello, World!")

	// Prints the current time using the Now function from the time package aliased as t.
	fmt.Println(t.Now())
}
