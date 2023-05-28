package main

import (
	"fmt"
)
// previous similar syntax
// type deck []string

type contact struct {
	email string
	phone string
}

// type person which is type struct, a simple object
type person struct {
	firstName string
	lastName string
	contact
}

// RAM ALLOCATION

// Address		Value
// 0000			  -
// 0001			  In main function it creates the first user struct
// 0002			  -
// 0003			  func (p person) updateName(newFirstName string) | here p is the copy of person with different memory allocation
// 0004			  -



func main(){
	user := person{
		firstName:"Sankar",
		lastName:"Prasad Biswas",
		contact: contact{
			email: "sankarbiswas07@gmail.com",
			phone: "+91 8961766682",
		},
	}

	userPointer := &user // assigning memory address

	userPointer.updateName("S.")
	user.print()
}

// RECEIVER FUNCTION WITH STRUCTS

func (p person) print(){
	fmt.Printf("%+v \n", p)
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}