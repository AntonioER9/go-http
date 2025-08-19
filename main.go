package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
}
