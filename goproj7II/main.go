package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// This is how we create a new channel
	// The 'make' built-in function that will create a value of the same type
	// The 'c' function, can only be used in the same function it was created
	// We need to make sure we pass this channel into the checkLink function
	// Then checkLink, which will be executed inside of a go routine will
	// have the ability to communicate back over to this main function
	c := make(chan string)

	for _, link := range links {
		// We will call our channel in as an arugement to the checkLink function
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

// Then we will also pass the channel in as a arugement to the checkLink func
// To do this, you will, say 'channel' then the type the channel is of
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		// This means the send a message into our channel
		c <- " might be down"
		return
	}
	fmt.Println(link + " is up!")
	c <- " It's up"
}

// The reason when you run this function with the 'go' keyword, you will
// have no output because the main go routine gets done before the child
// go routines exeuctue their code, making a empty output

// The way to fix this is by using go channels. We are going to use a
// channel to make sure the main routine is aware of when each of the
// child go routine completes their code. Channels are the only way
// to communicate with go routines

// The most important thing to know about channels is that they are typed
// The information that we pass into a channel, or the data that we share
// between the different routines must all be of the same type
// Example:
// Go routine -> Channel of type string (using string) -> go routine = GOOD
// Go routine -> Channel of type string (using int) -> go routine = BAD

// SENDING DATA WITH CHANNELS:
// channel <- 5
// Send the value '5' into this channel

// myNumber <- channel
// Wait for a value to be sent into the channel. When we get one, assign
// the value to 'myNumber'

// fmt.Println(<- channel)
// Wait for a value to be sent into the channel. When we get one,
// log it out immediately
