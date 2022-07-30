package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// When making a request to a url endpoint, you have to make the 'http://URL_NAME
	resp, err := http.Get("http://google.com")
	// var nil Type || Type must be a pointer, channel, func, interface, map, or slice type
	if err != nil {
		fmt.Println("Error:", err)
		// func Exit(code int) Exit causes the current program to exit with the given status code.
		// Conventionally, code zero indicates success, non-zero an error.
		os.Exit(1)
	}
	fmt.Println(resp)
}

// Some info around the response struct in 'func (c *Client) Head(url string) (resp *Response, err error)'
// When it comes to grabbing the body, it will be of type 'io.ReadCloser'
// EX: Response struct:
// Status -> String
// StatusCode -> int
// Body -> io.ReadCloser

// io.ReadCloser Interface:
// Reader -> (io.Reader Interface) Read([]byte) (int, error)
// Closer -> (io.Closer Interface) Close() (error)

// NOTE: In go, we can take multiple interfaces and
// assemble them together to make another interface
// Ex:
// type ReadCloser interface {
// 	Reader
// 	Closer
// }
