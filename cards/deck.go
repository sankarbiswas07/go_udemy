package main

import "fmt"

type deck []string

// here "d" is the receiver
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	 }
}