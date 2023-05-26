package main

import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// In Go: there are two types of array - 1. fixed size & 2. slice(fancy name of a array which can grow or shrink)
	// here we are writing slice and iterate over it

	cards := [] string { "Ace of Heart", newCard()}
	cards = append(cards, "Four of Heart")

	for i, card := range cards {
	 fmt.Println(i, card)
	}
}
