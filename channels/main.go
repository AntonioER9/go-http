package main

import "net/http"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "is down")
		return
	}
	println(link, "is up")
}
