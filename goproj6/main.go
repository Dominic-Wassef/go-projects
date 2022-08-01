package main

import (
	"fmt"
	"io"
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
	// The 'make' function is a built in function in the language that takes a type of a slice
	// And as a second arugement, '99999' this is the number of elements or empty spaces
	// we want the slice to be initalized with.

	// The reason we gave such a large int arugement
	// to the read function is because it's not set up to resize
	// the slice if the slice if already full, so we give it a
	// large number to avoid that byte slice being too small
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// There is a way to condense the above code
	// without using a byte using the writer interface
	// Difference between Reader and Writer:
	// Reader:
	// Source of data -> Reader -> []byte
	// []byte -> writer -> Some form of output (Ex: Outgoing HTTP Request, Terminal, etc)
	// To make it work we need to find something in the standard library that implements
	// the Writer interface, and use that to log out all the data that we're recieving
	// from the reader

	// func Copy:
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// The first arugment is some value that implements the writer interface
	// The second arugment is something that implements the reader interface

	// os.Stdout implements the writer interface which is a value that is export by os
	// The Stdout is of type file, file is a type inside of go that implements the write interface

	io.Copy(os.Stdout, resp.Body)
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

// The io.Reader interface:
// Rather than having all of these seperate types, which would mean we would have
// to write different functions to read from all the different sources of data
// EX:
// Source of input:
// HTTP Request Body -> Reader -> []byte (Output data that anyone can work with)
// Text file on a hard drive -> Reader -> []byte
