// Conditions in Go:
// - Comparison operators return true or false when comparing two values
// - if/else if/else statements are evaluated from top to bottom
// - Only the block of the first true condition is executed
// - An else if statement must directly follow an if statement

// Switch statements in Go:
// - A switch statement simplifies multiple conditions compared to if statements
// - The switch keyword is followed by a value to compare
// - Each case checks if the value matches and executes its block if true
// - When a case matches no other cases are evaluated
// - The default block runs if no case matches

// Variable scope in blocks:
// - Variables declared in an if or switch block are scoped to that block
// - Variables scoped to a block cannot be accessed outside it

// Random numbers and seeding in Go:
// - Random numbers need a seed to appear random
// - Use 'rand.Seed()' to set the starting point for random numbers
// - A constant seed generates the same sequence of numbers every run
// - Use 'time.Now().UnixNano()' to create a unique seed for randomness
// - A unique seed ensures random numbers differ with every program run

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var bitcoinPrice int = rand.Intn(100000)
	var ethereumPrice int = rand.Intn(10000)

	if bitcoinPrice >= 90000 {
		fmt.Printf("Decision: SELL Bitcoin ($%d).\n", bitcoinPrice)
	} else if bitcoinPrice >= 50000 {
		fmt.Printf("Decision: HOLD Bitcoin ($%d).\n", bitcoinPrice)
	} else {
		fmt.Printf("Decision: BUY Bitcoin ($%d).\n", bitcoinPrice)
	}

	switch {
	case ethereumPrice > 9000:
		fmt.Printf("Decision: SELL Ethereum ($%d).\n", ethereumPrice)
	case ethereumPrice >= 8000 && ethereumPrice <= 9000:
		fmt.Printf("Decision: HOLD Ethereum ($%d).\n", ethereumPrice)
	default:
		fmt.Printf("Decision: BUY Ethereum ($%d).\n", ethereumPrice)
	}
}
