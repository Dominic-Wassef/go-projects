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
	// When using go channels, you have to match the amount you want that channel
	// to print with the amount of data. Here we are going to use the for-loop

	// This will call the fmt.Println function and recieve a value over our channel
	// a number of times equal to the number of strings we have in our slice

	// We will not be using a range helper, We will iteriate from 0 up to 1 - number
	// of strings in slice.

	// We use the for-loop because the value of (<-c) is a blocking cal so we need
	// to use it in order to iterate through our data held in the links []string
	// Note the for-loop isn't exeucted at once, instead it iterates each one:

	// for i := 0; i < len(links); i++ {
	// 	(<-c)
	// }

	// Here, We are not going to ping each of the website just one time, we want to
	// ping them over and over again. Every time we make a request to a url, once
	// that routine finishes up, we will want to ping them again

	// Know that the first arugement (<-c) can be of type string because we know
	// that we will be recieving a string

	// With this for-loop we can use the range helper with a channel
	// We are going to run through the for-loop every time a channel emits value
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Below we are replacing checkLink func with a function literal
	// In go, a function literal is some unnamed function that we use to wrap
	// some little chunk of code so we can execute some point in the future
	for l := range c {
		// In function literal, we put the go keyword in front of the word func
		go func(l string) {
			// time.Sleep (standard library) pauses the current goroutine
			time.Sleep(10 * time.Second)
			// loop variable l captured by func literal:
			// Right now inside of our main routine whenever we recieve a value from
			// our channel, we are assigning it to the variable l

			// ERR: The function literal is referencing l, it is inside the outer-scope
			// of the function literal. This means that the function literal is being
			// exeucted in a seperate go routine. You are referencing a variable thats
			// declared in the out scope of this function literal. Because of this,
			// the child routine will infintely call only 1 url address

			// Remember, we never try to access the same variable from a different go
			// routine. We share info with a new go routine by passing in an arugement
			// To fix this, we will provide l as an arugement to the function literal
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Has an error!", err)
		c <- link
	}
	fmt.Println(link, "Ping!")
	c <- link
}
