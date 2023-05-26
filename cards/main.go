package main

import "fmt"

func main() {
	// var card string = "Ace of Spread"
	// the above line is equivalent to the following line

	// := telling go , to figure ut what type of data we are assigning to it
	// in case of the following. it will set a string type
	// := only valid while the initialization, no re-init state.

	card := "Ace of Spread"

	// Re-assign with equal to operator
	card = "Fives diamonds"

	fmt.Println(card)
}
