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

	fmt.Println(<-c)
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
	return
}
