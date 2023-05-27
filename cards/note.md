

// go run main.go deck.go
// Here we have made a custom dec type which is actually a slice of type string
// type deck []string

======================

// Making a function which will make the cards list
// when we do not use index variable, put _ : go will understand and not showing errors
// func newDeck()

======================

// Now we need a deal function which will return the cards which has been distributed after n number of deal
// and the remaining cards.
// Remember, both the return types are of deck
// func deal(d deck, handSize int) (deck, deck)

======================

// here "d" is the receiver
// Actual copy of the deck we'er working with is d
// (d deck) deck means, every deck type variable will have print() method
// d can be cards also. but by convention it is the first letter or first a few latter as variable name
// in go, it good practice ot have first letter d from deck name

// like OO approach
// deck class will have a deck instance with will have print() method
// go is not a Object oriented language.
// we can work in similar way by writing Receiver function bellow-
// func (d deck) print()

======================

// assign multiple variables form multiple return value
// hand, remainingDeck := deal(cards, 5)
// hand.print()
// remainingDeck.print()

======================

//  PLAN : CARDS === newDeck > print > shuffle > deal > saveToFile > newDeckFromFile

// from the above plan, we are making small parts. 
// 1. Make custom type of deck which is a slice.
// 2. Deck creation function newDeck().
// 3. We have to save it in a file. so we need type conversion. file write needs Byte type slice