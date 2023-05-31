package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Print((eb.getGreeting()))
}

// func printGreeting(sb spanishBot) {
// 	fmt.Print((sb.getGreeting()))
// }

func (eb englishBot) getGreeting() string {
	// different very much different custom logic
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// NOTE : 54. Problems Without Interfaces
// The above code block is not allowing two greeting function
// Functions with same name is not allowed in the same file
