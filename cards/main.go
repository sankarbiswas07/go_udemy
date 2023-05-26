package main

// import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// In Go: there are two types of array - 1. fixed size & 2. slice(fancy name of a array which can grow or shrink)
	// here we are writing slice and iterate over it

	cards := deck { "Ace of Heart", newCard()}
	cards = append(cards, "Four of Heart")

	cards.print()

}


// go run main.go deck.go
// Here we have made a custom dec type which is actually a slice of type string