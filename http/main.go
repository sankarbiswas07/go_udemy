package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	lw := logWriter{}

	// Short-hand
	// se io.Copy
	io.Copy(lw, resp.Body)

	bs := make([]byte, 99)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
	fmt.Println(resp)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just write this many bytes:", len(bs))
	return len(bs), nil
}
