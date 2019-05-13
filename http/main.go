package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom Writer type
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Body is a byte slice
	// bs := make([]byte, 99999)

	// Pass byte slice to the Read function
	// resp.Body.Read(bs)

	// Custom Writer implemtation
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
