package main

import (
	"fmt"
)

type colors map[string]string

func main(){
	colorSet := colors{
		"red": "#FF1234",
		"green": "#FF1234",
		"white": "#FF1234",
	}

	fmt.Println("\nfirst time assign - way one : you can see bellow syntax to get other ways too")
	fmt.Println(colorSet)
	
	fmt.Println("Red changed by re-assign")
	colorSet["red"] = "123456"
	fmt.Println(colorSet)
	
	fmt.Println("green changed by receiver function")
	colorSet.updateColor("green", "123457")
	
	fmt.Println("New Key: yellow added\n")
	colorSet["yellow"] = "QWERTY"	

	fmt.Println("Print Color set via a function\n")
	printMap(colorSet)
	
	fmt.Println("Delete key red\n")
	delete(colorSet, "red")

	fmt.Println("Print Color set via a function\n")
	printMap(colorSet)
}

func (c colors) updateColor(whichColor string, colorCode string) {
	c[whichColor] = colorCode
}

func printMap (c colors){
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

// NOTE : Diff between Map & Struct

// MAP

// 1. All keys much be the same type 
// 2. All values must be the same type
// 3. Keys are indexed- we can iterate over them
// 4. Use to represent a collection  of related properties
// 5. Don't need to know all the keys at compile time
// 6. Reference Type!


// Struct:
// 1. Values can be of different type
// 2. Keys don't support indexing
// 3. You need to know all the different fields at compile time
// 4. Use to represent a "thing" with a lot of different properties
// 5. Value Type!