Structs:
  - similar to Simple Object in JavaScript.
  - Organize data with Struct

when

// type person which is type struct, a simple object
type person struct {
	firstName string
	lastName string
}


	// // WAY:1 - to create struct, arguments will create struct in a sequential order.
	// // firstName "Prasad Biswas" && lastName "Sankar" IF fullName_1 := person{"Prasad Biswas", "Sankar"}
	// fullName_1 := person{"Sankar", "Prasad Biswas"}
	// fmt.Println(fullName_1)

	// // WAY:2 - to create struct, arguments will create struct by mentioning key.
	// fullName_2 := person{firstName:"Sankar", lastName:"Prasad Biswas"}
	// fmt.Println(fullName_2)

	// // WAY:3 - to create struct, create a variable of type person
	// // this actually means there is all the keys with "" empty string
	// var fullName_3 person
	// fmt.Println(fullName_3)
	// fmt.Printf("%+v \n", fullName_3)
	// // UPDATE
	// fullName_3.firstName = "Sankar"
	// fullName_3.lastName = "Prasad Biswas"
	// fmt.Println(fullName_3)
	// fmt.Printf("%+v \n", fullName_3)

	// 41. Embedding Structs
	// This is nothing but nested objects
	// step 1: like person, you have to create one struct "type"
	// step 2: make a key in the parent object of that particular struct "type"(step 1)



  =========================

  POINTER: 
  // RAM ALLOCATION

// Address		Value
// 0000			  -
// 0001			  In main function it creates the first user struct
// 0002			  -
// 0003			  func (p person) updateName(newFirstName string) | here p is the copy of person with different memory allocation
// 0004			  -


// KEY NOTE POINT -
// Turn "Value" into "Address" with "&value"
// Turn "Address" into "Value" with "*address"