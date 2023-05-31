package main

import "fmt"

// Our program has a new typw called "bot"

// we use interface  -
// to define a method set or function set. what
// different function it has and wht return type it has
type bot interface {
	// Now,
	// fro all the custom type, inside the definition
	// if you are a type in this program with a
	// function called "getGreeting" and you return a
	// string then you are now an honorary member
	// of the type bot

	// no that you're(englishBot, spanishBot) [struct]
	// also an honorary member of type "bot". eg: -
	// type englishBot struct{}
	// type spanishBot struct{}
	// you can now call this function called "printGreeting"

	// in simple way
	//

	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// As printGreeting(eb) and printGreeting(sb) are different type.
// and they are honored member as bot
// So, instead of having different type.
// printGreeting will allow to let them in side the function
//
//

func printGreeting(b bot) {
	fmt.Println((b.getGreeting()))
}

// func printGreeting(eb englishBot) {
// 	fmt.Print((eb.getGreeting()))
// }

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
