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
	c := make(chan string)
	people := [5]string{"Baby", "Tiger", "8luebottle", "Blue", "She"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c) //  Loop : len(people)
	}

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

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is Sexy" // Send "true" to the Channel c
}
