package main

import (
	"fmt"
	"net/http"
)

func main() {
	ls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://hulu.com",
		"http://stackoverflow.com",
	}
	for _, l := range ls {
		d(l)
	}
}

func d(s string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s + " is down")
		return
	}
	fmt.Println(s + " is up")
}
