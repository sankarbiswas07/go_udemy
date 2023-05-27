package main

// import "fmt"

func main() {

	cards:= newDeck()

	cards.saveToFile("my_deck")

	// fmt.Println(cards.toString())

}