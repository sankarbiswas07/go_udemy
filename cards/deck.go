package main

import (
	"fmt"
	"strings"
	"os"
	"math/rand"
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

func (d deck) shuffle() deck {
	rand.Shuffle(len(d), func(i, j int){
		d[i], d[j] = d[j], d[i]
	})
	return d
}


func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// byteString
	bs, err := os.ReadFile(filename)
	if(err != nil){
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	s:= strings.Split(string(bs),",")
	return deck(s)
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