package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request Failed!")

func main() {
	// Testing Go Channel
	c := make(chan bool)
	people := [2]string{"Baby", "Tiger"}
	for _, person := range people {
		go isSexy(person, c)
	}
	result := <-c // Receive it from isSexy
	fmt.Println(result)

	var results = make(map[string]string)
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
	results["hello"] = "Hello"
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

// Check one by one --> checking like Python
// Later, for concurrency I'm going to use goroutine
func hitURL(url string) error {
	fmt.Println("Checking URL: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true // Send "true" to the Channel c
}
