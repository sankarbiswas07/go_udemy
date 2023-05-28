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


func main(){

	fullName_4 := person{
		firstName:"Sankar",
		lastName:"Prasad Biswas",
		contact: contact{
			email: "sankarbiswas07@gmail.com",
			phone: "+91 8961766682",
		},
	}
	fmt.Println(fullName_4)
	fmt.Printf("%+v \n", fullName_4)

}