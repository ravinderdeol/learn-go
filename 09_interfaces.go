// Interfaces:
// - Interfaces define methods that types can implement.
// - A type implements an interface by providing all the required methods.
// - Interfaces allow flexible and reusable code without explicit declarations.

// Type Assertions and Switches:
// - Type assertions check the concrete type of an interface using x.(T).
// - Safe assertions return a value and success flag to avoid panics.
// - Type switches handle multiple types in an interface safely.

// Empty Interface:
// - The empty interface (interface{} or any) can hold a value of any type.
// - Use type assertions or switches to retrieve values from an empty interface.
// - Avoid using empty interfaces excessively to ensure type safety.

// Best Practices for Interfaces:
// - Keep interfaces small with one or two methods for simplicity.
// - Define interfaces based on behavior to improve flexibility and maintenance.

// Pointers:
// - Dereferencing (*ptr) gets or changes the value at a memory address.
// - Use the & operator to pass a variable's address to a function.

package main

import "fmt"

// Vehicle is an interface that requires the CalculateRentalCost method.
type Vehicle interface {
	CalculateRentalCost() int
}

// PrintRentalCost accepts any type that implements the Vehicle interface.
func PrintRentalCost(v Vehicle) {
	fmt.Printf("Rental Cost: $%d\n", v.CalculateRentalCost())
}

// Scooter is a struct that represents a scooter rental with hours of usage.
type Scooter struct {
	hours int
}

// Car is a struct that represents a car rental with hours and miles of usage.
type Car struct {
	hours int
	miles int
}

// CalculateRentalCost calculates the cost of renting a scooter based on hours.
func (s Scooter) CalculateRentalCost() int {
	const hourlyRate = 5 // Rental cost per hour for a scooter.
	return s.hours * hourlyRate
}

// CalculateRentalCost calculates the cost of renting a car based on hours and miles.
func (c Car) CalculateRentalCost() int {
	const hourlyRate = 10 // Rental cost per hour for a car.
	const mileageRate = 2 // Rental cost per mile for a car.
	return (c.hours * hourlyRate) + (c.miles * mileageRate)
}

func main() {
	// Create a Scooter instance with 3 hours of usage.
	scooter := Scooter{hours: 3}

	// Create a Car instance with 2 hours of usage and 50 miles driven.
	car := Car{hours: 2, miles: 50}

	// Print the rental cost of the scooter using the interface.
	PrintRentalCost(scooter)

	// Print the rental cost of the car using the interface.
	PrintRentalCost(car)
}
