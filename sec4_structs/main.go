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

	// WAY:1 - to create struct, arguments will create struct in a sequential order.
	// firstName "Prasad Biswas" && lastName "Sankar" IF fullName_1 := person{"Prasad Biswas", "Sankar"}
	
	fullName_1 := person{"Sankar", "Prasad Biswas"}
	fmt.Println(fullName_1)

	// WAY:2 - to create struct, arguments will create struct by mentioning key.
	fullName_2 := person{firstName:"Sankar", lastName:"Prasad Biswas"}
	fmt.Println(fullName_2)

}