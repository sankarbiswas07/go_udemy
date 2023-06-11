package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.org",
	}

	// creating a channel
	// c is scoped and available inside main()
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//76. Receiving message

	// message is a blocking operating
	// when for ends. the code block wait until a message comes in channel
	// channel closed immediately when first message come
	// Everything in a main.go will be in a main routine

	// simply listen in channel means blocking till come first message
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// if you uncomment it. terminal will wait for signal
	// as <- means a blocking, in such case next "END" will never be printed.

	//77. Receiving messages | one way
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// 78. Repeating Routines
	// for{
	// 	go checkLink(<-c, c)
	// }

	// 79. Alternative Loop syntax
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// 80,81,82 | a good way to implement a timer control
	for l:= range c {
		go func(link string){
			time.Sleep(3000 * time.Millisecond)
			checkLink(link, c)
		}(l)
	}

	fmt.Println("END")
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !")
		c <- link
		return
	}
	// How can we communicate or using channel
	fmt.Println(link, "is up!")
	c <- link
}
