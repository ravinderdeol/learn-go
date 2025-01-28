// How to compile a file:
// - Use 'go build <filename>.go' to create an executable
// - Run the executable using './<filename>'

// Compiler:
// - Converts Go code into a standalone executable
// - Allows running without Go installed on the system
// - Packages all dependencies into the binary
// - Generates optimized machine code for better performance

// Compiling and running programs:
// - Use 'go run <filename>.go' for development or testing
// - Use 'go build <filename>.go' to create a distributable executable

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
