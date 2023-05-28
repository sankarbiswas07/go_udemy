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


// 47. Gotchas With Pointer:
// Go is a Pass By Value type language
// Either you can Pass value as functional argument -  when we
// Or you can make Receiver function
// You can use the copy of the actual value or you can pass the pointer 

// Conclude - we can pass slice in function and modify directly with ut pointer, But with Struct - it is working differently (gotcha !!!)


// 48. Reference vs Value types

// Slices are a type of array which can grow and shrink
// Slice is data structure that records the length of the slice(current), the capacity of the slice, and a reference to the underlying array
// Slice data structure has it's own structure and a fixed size array structure underlying data structure

===========================================================
// Value Types                          Reference Types
// Use pointers to change               Don't worry about
// these things in a func               pointer with these
===========================================================
// int                                  slices
// float                                maps
// string                               channels
// bool                                 pointers
// structs                              functions
===========================================================