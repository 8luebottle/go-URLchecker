package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed!")

func main() {
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.apple.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"http://babytiger.netlify.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
}

// Beautiful Go's Concurrency
func hitURL(url string, c chan<- result) { // Channel that send only
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
