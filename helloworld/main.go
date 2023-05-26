// Hello world type program

package main

/**
* QUESTIONS? => Answer
*
*	1. How do we run the code in our project? => go command line tool. either we go to project directory (helloworld) > go run main.go or we can build (go build main.go) it and ./ main
* 2. What does "package main" mean? => if we have multiple file and relates each other. wrap it within main. package main
* 3. What does "import 'fmt'" mean? => default the package main has no connection with any of the packages. "fmt" is the standard lib of go
* 4. What is "func" thing? => func is keyword for function
* 5. How is the main.go file organized?
*
 */

import "fmt"

func main() {
	fmt.Println("Hi there !")
}
