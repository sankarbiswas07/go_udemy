package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Short-hand
	// se io.Copy

	io.Copy(os.Stdout, resp.Body)


	// bs := make([]byte, 99)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))
	// fmt.Println(resp)
}
