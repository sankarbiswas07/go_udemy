package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// What is the type of response ?
	// https://pkg.go.dev/net/http@go1.20.4#Response

	// Response's main thing is the body part,
	// Which is type Body io.ReadCloser,
	// let's get into thing type

	// check ReadMe.md for
	// understating of Response => https://pkg.go.dev/net/http@go1.20.4#Response
	// So we can understand the response.body type "io.ReadCloser"
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
