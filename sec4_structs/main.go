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


// KEY NOTE POINT 1 -
// Turn "Value" into "Address" with "&value"
// Turn "Address" into "Value" with "*address"

// KEY NOTE POINT 2 -
// Difference between *person & *pointerToPerson
// *person : in receiver func (pointerToPerson *person), this is an "type description" - it means we're working with a pointer to a person
// *pointerToPerson : in (*pointerToPerson).firstName, this is an "operator" - it means we want to manipulate the value the pointer is referring

func main(){
	user := person{
		firstName:"Sankar",
		lastName:"Prasad Biswas",
		contact: contact{
			email: "sankarbiswas07@gmail.com",
			phone: "+91 8961766682",
		},
	}

	userPointer := &user // assigning memory address, & = returns a pointer, a memory address access not the struct directly

	userPointer.updateName("S.")
	user.print()
}

// RECEIVER FUNCTION WITH STRUCTS

func (p person) print(){
	fmt.Printf("%+v \n", p)
}

// *person => not the actual value of the pointer(person) pointed to a memory address
// it is type description not the operator to get the value
// conclude : userPointer is equal to pointerToPerson which is referring the pointer of user struct
func (pointerToPerson *person) updateName(newFirstName string){

	// *pointerToPerson means take the pointerToPerson and give the value of the pointer
	(*pointerToPerson).firstName = newFirstName // So we can able to update the firstName of the value of user
}