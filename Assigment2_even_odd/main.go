package main

import (
	"fmt"
)


func main()  {

	slice := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, n := range slice {
		whatType := "even"
		if n % 2 == 0 {
			whatType = "even"
		}else{
			whatType = "odd"
		}
		fmt.Printf("%v is %v \n", n, whatType)
	}
}