// ASSIGNMENT:

// Write a program that creates two custom struct types called
// 'triangle' and 'square'

// The 'square' type should be a struct with a 
// field called 'sideLength of type float64

// The 'triangle' type should be a struct with a field 
// called 'height' of type float64 and a field of type base' of type float64

// Both types should have function called 'getArea' 
// that returns the calculated area of the square or triangle

// Area of a triangle = 0.5 * base * height.
// Area of a square - sideLength * sideLength

// Add a 'shape' interface that defines a function called 'printArea'.
// This function should calculate the area of the given shape and 
// print it out to the terminal Design the interface so that 
// the 'printArea' function can be called with either a triangle or a square.

// SOLUTION:

package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base float64
}

func main(){
	// fmt.Println("Assignment 3 : Interface")
	tr := triangle{
		base: 3,
		height: 5,
	}
	sq := square{
		sideLength: 4,
	}

	fmt.Println("Area of the triangle", printArea(tr))
	fmt.Println("Area of the square", printArea(sq))
}

func printArea(sh shape) float64{
	return sh.getArea()
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return math.Pow(sq.sideLength, 2) // sideLength^2
}