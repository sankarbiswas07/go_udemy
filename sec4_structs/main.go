package main

import (
	"fmt"
)
// previous similar syntax
// type deck []string

// type person which is type struct, a simple object
type person struct {
	firstName string
	lastName string
}

func main(){

	// one wat to create struct, arguments will create struct.
	// firstName "Prasad Biswas" && lastName "Sankar" IF fullName_1 := person{"Prasad Biswas", "Sankar"}
	fullName_1 := person{"Sankar", "Prasad Biswas"}
	fmt.Println(fullName_1)
}