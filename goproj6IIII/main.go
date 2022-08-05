package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	lin, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, lin)
}
