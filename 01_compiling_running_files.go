// How to compile a file:
// - Enter go build <filename>.go in the terminal.
// - This creates an executable file in the same directory.
// - Enter ./<filename> in the terminal to run the executable.

// Compiler:
// - Converts Go code into a standalone executable file.
// - The executable runs without requiring Go to be installed.
// - Packages all dependencies into the binary for easy sharing.
// - Produces optimized machine code for fast execution.

// Compiling and running programs:
// - Use 'go run' for development and testing.
// - Use 'go build' to create an executable for distribution.

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
