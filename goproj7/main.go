package main

import (
	"fmt"
	"net/http"
)

// We will be using goRoutines to make indiviual requests to the url
// and be able to print the status and move on to the next link
// When we lauch a go program, we will lauch one goRoutine
func main() {
	links := []string{
		"http://google.com",
	}
	for _, link := range links {
		// Control flow is passed back to our main go routine
		// It will then say it looks like this go routine can't do anything
		// so the second go routine will run and make the get request to url
		// the 'go' keyword in front will create a new thread go routine
		// and checkLink(link) will run this function with the go routine
		checkLink(link)
	}
}

func checkLink(link string) {
	// This is the blocking call which means that whle this function is
	// being executed, the main go route can do nothing else

	// We will fix this problem by lauching mulitple goroutines
	// When we call the check link function, we will place the keyword
	// 'go' in front of it which means to run this function inside of a
	// brand new go routine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		return
	}
	fmt.Println(link + " is up!")
}

// IMPORTANT INFO:
// The GO Scheduler runs one routine until it finishes or
// makes a blocking call (like an HTTP Request)

// One CPU Core -> Go Scheduler -> GO Routine I -> GO Routine II -> <-COPY
// Only one go routine runs at a time and once it's done, it will continue

// By default GO tries to use one core, although you can have it run
// multiple CPU Core, which makes it where you can run mutilple go routines
// at once. Remember that concurrency is not parrellism with go routine

// Concurrency - We can have multiple threads executing code. If one
// thread blocks, another one is picked up and worked on

// Parallelism - Multiple threads executed at the exact same time.
// This requires multiple CPU Core's

// Our Running Program:
// Main Routine - Main routine created when we lauched our program
// Child go routine  - Child routines created by the 'go' keyword
