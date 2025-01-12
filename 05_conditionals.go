// Conditions:
// - Comparison operators compare two values and return true or false.
// - if/else if/else statements are checked from top to bottom.
// - Only the first true condition runs its block of code.
// - The else if statement must follow an if statement.

// Switch:
// - A switch statement is simpler than multiple if statements.
// - The switch keyword is followed by a value to compare.
// - Each case checks if the value matches the case's value.
// - If a match is found the case's block runs and no other cases are checked.
// - The default statement runs if no cases match.

// Variable scope:
// - Variables declared in if or switch statements are scoped to those blocks.
// - Variables scoped to these blocks cannot be accessed outside them.

// Random numbers and seeding:
// - Random numbers in Go need a unique seed to be random.
// - The rand.Seed() method sets the starting point for random numbers.
// - Passing a constant to rand.Seed() generates the same numbers every run.
// - Use time.Now().UnixNano() to create a unique seed.
// - A unique seed ensures different numbers every time the program runs.

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
