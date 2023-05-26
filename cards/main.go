package main


func main() {

	cards:= newDeck()
	cards.print()

}


// go run main.go deck.go
// Here we have made a custom dec type which is actually a slice of type string