package main

import "fmt"

func main() {

	cards:= newDeck()

	fmt.Println(cards.toString())

	// assign multiple variables form multiple return value

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

}


// go run main.go deck.go
// Here we have made a custom dec type which is actually a slice of type string