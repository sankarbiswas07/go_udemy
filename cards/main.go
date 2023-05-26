package main

import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	card := newCard()

	fmt.Println(card)
}
