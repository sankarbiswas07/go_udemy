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