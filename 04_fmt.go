// Print functions in Go:
// - 'fmt.Print()' outputs arguments without spaces or a line break
// - 'fmt.Println()' outputs arguments with spaces and adds a line break
// - 'fmt.Sprint()' returns concatenated values without spaces
// - 'fmt.Sprintln()' returns concatenated values with spaces
// - 'fmt.Printf()' formats strings using placeholders and prints the result
// - 'fmt.Sprintf()' formats strings using placeholders and returns the result
// - 'fmt.Scan()' reads user input and stores it in variables

// Verbs for formatting:
// - Verbs start with '%' followed by a letter (e.g. '%v' for value '%T' for type '%.2f' for a float with 2 decimals)

// Input handling with 'fmt.Scan()':
// - Variables used with 'fmt.Scan()' require the '&' operator to pass their address (e.g. 'fmt.Scan(&tickerSymbol)')
// - 'fmt.Scan()' stops reading input at the first space or newline

package main

import "fmt"

func main() {
	var tickerSymbol string
	var tickerPrice float32

	fmt.Println("Enter the ticker symbol: ")

	fmt.Scan(&tickerSymbol)

	fmt.Printf("Enter the ticker price (%s): ", tickerSymbol)
	fmt.Scan(&tickerPrice)

	var output string = fmt.Sprintf("The price of %s is $%.2f.", tickerSymbol, tickerPrice)
	fmt.Println(output)
}
