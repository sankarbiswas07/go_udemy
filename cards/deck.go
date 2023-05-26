package main

import "fmt"

type deck []string

// Making a function which will make the cards list
func newDeck() deck {
	cards := deck{}

	cardSuites:=[] string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues:=[] string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _,suit := range cardSuites{
		for _,value := range cardValues{
			cards = append(cards, value+" of "+suit)		
		}
	}

	return cards
}

// here "d" is the receiver
// Actual copy of the deck we'er working with is d
// (d deck) deck means, every deck type variable will have print() method
// d can be cards also. but by convention it is the first letter or first a few latter as variable name
// in go, it good practice ot have first letter d from deck name

// like OO approach
// deck class will have a deck instance with will have print() method
// go is not a Object oriented language.
// we can work in similar way by writing Receiver function bellow-
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	 }
}