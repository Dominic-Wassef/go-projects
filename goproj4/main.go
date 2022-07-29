package main

import "fmt"

func main() {
	// A map is the equalivent of objects in javascript
	// 	When declaring a empty map, it will log its 0 value, meaning and empty map
	// Using var is good for empty maps
	// var colors map[string]string
	// ^ is equalvant to the bottom example
	// colors := make(map[int]string)

	// map[string]string{} is a map with string keys and string values
	// map[string]int{} is a map with string keys and int values
	colors := map[string]string{
		"green": "4bf745",
		"white": "ffffff",
		"red":   "ff0000",
	}

	// Notice how we dont use the appendation like how we did with structs:
	// EX structName.white wouldn't work with maps, instead we always we mapName[0]
	// colors[10] = "ffffff"

	// We can delete keys and value off of a map using the delete() function
	// delete(colors, 10)

	printMap(colors)
}

// We can iterate over the map like so:
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Some differences between structs and maps

// Map: All keys AND values must be the same type
// Key are indexed meaning that we can iteriate over them
// Referenced Type! (HINT: Pointers won't have to be used here)
// With a map, you can change any value or delete it because not all keys at compile

// Struct: Values can be of any type
// Keys don't support indexing meaning we can't iteriate over them
// Value Type! (HINT: Pointers will have to be used here)
// With a struct, you must know all keys and defined them at compile time
