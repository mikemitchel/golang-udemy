package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// pass var l value as argument between routines
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // passed from outer scope as copy
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // blocking call

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
