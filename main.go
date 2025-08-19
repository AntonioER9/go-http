package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

// logWriter is a custom writer that logs data to the console.
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Print(string(bs))
	return len(bs), nil
}
