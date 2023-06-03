package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}


	// Short-hand
	// se io.Copy
	io.Copy(lw, resp.Body)


	// bs := make([]byte, 99)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))
	// fmt.Println(resp)
}

func (logWriter) Write(bs []byte) (int, error){
	return 1, nil
}
