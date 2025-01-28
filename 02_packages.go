// Code organization:
// - Go programs use packages to group related code
// - A package is defined by the 'package' declaration in a '.go' file
// - Files with the same 'package' declaration belong to the same package
// - A folder typically represents a package and contains its files

// Packages:
// - A package organizes and reuses code in Go
// - The 'package' declaration specifies the package a file belongs to
// - The 'main' package defines the entry point of a Go application
// - Go programs with a 'main' function in the 'main' package compile into an executable
// - The executable starts by running the 'main' function

// Documentation tools:
// - Use 'go doc' to view documentation for Go packages or functions
// - Use 'go doc <package>' to see details about a package (e.g. 'go doc time')
// - Use 'go doc <package>.<function>' to see details about a function (e.g. 'go doc time.Now')

// Imports in Go:
// - Imports enable code reuse from other packages
// - Use parentheses to import multiple packages in one statement
// - Use an alias to rename a package (e.g. 't' for 'time')

// Functions in Go:
// - Declare functions in Go using the 'func' keyword
// - A function name follows the 'func' keyword (e.g. 'main')
// - Function logic is enclosed within curly braces

package main

import (
	"fmt"
	t "time"
)

func main() {

	// Print 'Hello, World!' using 'fmt.Println'
	fmt.Println("Hello, World!")

	// Print the current time using 't.Now' from the 'time' package
	fmt.Println(t.Now())
}
