// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	content, err := os.ReadFile("fun.txt")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
//		os.Exit(1)
// 	}
// 	fmt.Println(string(content))
// }

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
