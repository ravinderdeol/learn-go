// Basic types in Go:
// - 'bool' is 1 bit and has values 'false' or 'true'
// - 'int8' is 8 bits with a range from -128 to 127
// - 'uint8' is 8 bits with a range from 0 to 255
// - 'int32' is 32 bits with a range from -2147483648 to 2147483647
// - 'uint64' is 64 bits with a range from 0 to 18446744073709551615

// Named and literal values:
// - A named value is a constant or variable with a name (e.g. 'const Pi = 3.14')
// - A literal value is the specific value assigned to a named value (e.g. '3.14')

// Constants in Go:
// - Use 'const' to define values that do not change (e.g. 'const maxSpeed = 70')
// - Constants remain fixed throughout the program execution
// - Use camelCase or PascalCase for descriptive constant names (e.g. 'maxSpeed')

// Variables in Go:
// - A variable is a named value that can change during program execution
// - Use 'var' to define variables with a type (e.g. 'var count int')
// - Variables do not need an initial value when declared (e.g. 'var name string')
// - Use ':=' to declare and assign a variable in one step (e.g. 'age := 30')
// - The type of a variable declared with ':=' is inferred from the assigned value (e.g. 'pi := 3.14' becomes 'float64')
// - Use operators like '+=', '-=', '*=', and '/=' for shorthand math operations (e.g. 'count += 10')
// - Declare multiple variables on a single line (e.g. 'var a, b int' or 'x, y := "Hello", 42')

// Default values for variables:
// - Numeric variables default to '0' before assignment
// - String variables default to an empty string ('')
// - Boolean variables default to 'false'

// Handling errors in Go:
// - The compiler stops creating a binary if an error is detected
// - Programs cannot run without a binary being created
// - Go provides the file name line number and column of the error (e.g. 'main.go:15:20')

package main

import "fmt"

func main() {
	const writer string = "Kazuo Ishiguro"
	const bookPublisher string = "Faber And Faber"

	var title string
	var year, pageNumber uint16
	var rating float32

	title = "The Remains Of The Day"
	year = 1989
	pageNumber = 258
	rating = 4.14

	fmt.Println("Title:", title)
	fmt.Println("Writer:", writer)
	fmt.Println("Publisher:", bookPublisher)
	fmt.Println("Year:", year)
	fmt.Println("Page Number:", pageNumber)
	fmt.Println("Rating:", rating)
	fmt.Println("")

	title = "Klara And The Sun"
	year = 2021
	pageNumber = 303
	rating = 3.74

	fmt.Println("Title:", title)
	fmt.Println("Writer:", writer)
	fmt.Println("Publisher:", bookPublisher)
	fmt.Println("Year:", year)
	fmt.Println("Page Number:", pageNumber)
	fmt.Println("Rating:", rating)
}
