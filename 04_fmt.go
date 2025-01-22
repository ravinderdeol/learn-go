// Print:
// - fmt.Print() prints arguments without spaces or a line break
// - fmt.Sprint() returns a value by concatenating arguments without spaces
// - fmt.Sprintln() returns a value by concatenating arguments with spaces
// - fmt.Println() prints arguments with spaces and adds a line break
// - fmt.Scan() gets input from the user and stores it in variables
// - fmt.Printf() interpolates strings using placeholders called verbs (supports multiple placeholders)
// - fmt.Sprintf() works like fmt.Printf() but returns the value instead of printing it

// Verbs:
// - Verbs are % followed by a letter (e.g. %v for value, %T for type, %.2f for float with 2 decimals)

// Input:
// - Variables passed to fmt.Scan() must use the & operator for their address (e.g. fmt.Scan(&tickerSymbol))
// - fmt.Scan() stops reading input at the first space or newline

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
