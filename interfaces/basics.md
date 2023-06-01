53. Purpose of Interface

We know that...

Every value has a type
Every function has to specify the type of its arguments


So does that mean...

Every function we ever write has to be rewritten to 
accommodate different types even if the logic in it is identical ?

For eg:

    func (d deck) shuffle() { ... some code ...} : in Assignment1_cards/deck.go : 23 : Can only shuffle a value of type "deck"

    Similarly,
        func (s []float64) shuffle() { ... some code ...} :  Can only shuffle a value of type "[]float64"
        
        func (s []string) shuffle() { ... some code ...} :  Can only shuffle a value of type "[]string"
        
        func (s []int) shuffle() { ... some code ...} :  Can only shuffle a value of type "[]int"

We actually can optimize the code a little by sing interfaces.
This is one of the main reason to implement interfaces.


==========================================================================

1. EnglishBot

type englishBot struct

// different very much different custom logic
func (englishBot) getGreeting() string

// holds identical/ general/ generic logic
//fmt.Println(eb.getGreeting()) // return "Hello there"
func printGreeting(eb englishBot)

1. SpanishBot

type spanishBot struct

// different very much different custom logic
func (spanishBot) getGreeting() string

// holds identical/ general/ generic logic
//fmt.Println(eb.getGreeting()) // return "Holla"
func printGreeting(eb spanishBot)



======================================

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
