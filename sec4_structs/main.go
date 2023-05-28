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

	// WAY:3 - to create struct, create a variable of type person
	// this actually means there is all the keys with "" empty string

	var fullName_3 person
	fmt.Println(fullName_3)
	fmt.Printf("%+v \n", fullName_3)
	
	// UPDATE
	fullName_3.firstName = "Sankar"
	fullName_3.lastName = "Prasad Biswas"

	fmt.Println(fullName_3)
	fmt.Printf("%+v \n", fullName_3)

}