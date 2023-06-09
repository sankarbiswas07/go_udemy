package main

import (
	"fmt"
	"net/http"
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

	// message is a blocking operating
	// when for ends. the code block wait until a message comes in channel
	// channel closed immediately when first message come
	// Everything in a main.go will be in a main routine

	// simply listen in channel means blocking till come first message
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c)
	// if you uncomment it. terminal will wait for signal
	// as <- means a blocking, in such case next "END" will never be printed.
	fmt.Println("END")
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "might be down !")
		c <- link + " might be down !"
		return
	}
	// How can we communicate or using channel
	// fmt.Println(link, "is up!")
	c <- link + " link is up !"
}
