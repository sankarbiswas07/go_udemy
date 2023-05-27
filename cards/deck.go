package main

import (
	"fmt"
	"strings"
	// "io"
	"io/ioutil"
)


type deck []string

// Receiver function
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	 }
}

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}

// ioutil.WriteFile is deprecated in latest go version
// so using io.WriteString -> https://pkg.go.dev/io@go1.20.4#WriteString

// func (d deck) saveToFile(filename string) error {
// 	return io.WriteString(filename, d.toString())
// }

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


// General function without parameter passing as an argument
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

// General function with parameter passing as an argument
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}