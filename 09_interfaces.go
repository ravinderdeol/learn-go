// Interfaces in Go:
// - Interfaces define a set of methods that a type can implement
// - A type implements an interface by providing all the required methods
// - Interfaces enable flexible and reusable code without explicit declarations

// Type assertions and switches in Go:
// - Type assertions check the concrete type of an interface using 'x.(T)'
// - Safe assertions return the value and a success flag to prevent panics
// - Type switches handle multiple types in an interface safely

// Empty interface in Go:
// - The empty interface ('interface{}' or 'any') can store a value of any type
// - Use type assertions or switches to extract values from an empty interface
// - Minimize empty interface usage to maintain type safety

// Best practices for interfaces in Go:
// - Keep interfaces small with one or two methods to ensure simplicity
// - Define interfaces based on behavior to improve flexibility and maintainability

// Pointers in Go:
// - Dereferencing ('*ptr') retrieves or modifies the value at a memory address
// - Use the '&' operator to pass a variable's address to a function

package main

import "fmt"

// Vehicle is an interface that requires the 'CalculateRentalCost' method
type Vehicle interface {
	CalculateRentalCost() int
}

// PrintRentalCost prints the rental cost of any type that implements the 'Vehicle' interface
func PrintRentalCost(v Vehicle) {
	fmt.Printf("Rental Cost: $%d\n", v.CalculateRentalCost())
}

// Scooter is a struct representing a scooter rental with hours of usage
type Scooter struct {
	hours int
}

// Car is a struct representing a car rental with hours and miles of usage
type Car struct {
	hours int
	miles int
}

// CalculateRentalCost calculates the rental cost of a scooter based on hours of usage
func (s Scooter) CalculateRentalCost() int {
	const hourlyRate = 5 // Cost per hour for a scooter
	return s.hours * hourlyRate
}

// CalculateRentalCost calculates the rental cost of a car based on hours and miles of usage
func (c Car) CalculateRentalCost() int {
	const hourlyRate = 10 // Cost per hour for a car
	const mileageRate = 2 // Cost per mile for a car
	return (c.hours * hourlyRate) + (c.miles * mileageRate)
}

func main() {
	// Create a Scooter instance with 3 hours of usage
	scooter := Scooter{hours: 3}

	// Create a Car instance with 2 hours of usage and 50 miles driven
	car := Car{hours: 2, miles: 50}

	// Print the rental cost of the scooter using 'PrintRentalCost'
	PrintRentalCost(scooter)

	// Print the rental cost of the car using 'PrintRentalCost'
	PrintRentalCost(car)
}
