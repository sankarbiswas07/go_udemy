package main

// import "fmt"

func main() {

	// cards:= newDeck()
	// cards.saveToFile("my_deck")

	cards := newDeckFromFile("my_deck")
	cards.print()

}