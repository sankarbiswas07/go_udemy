package main

import "fmt"

type deck []string


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

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	 }
}