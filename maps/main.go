package main

import "fmt"

// Summary: Maps vs Structs

// Use a map when you want to represent a collection of related properties.
// It's like a dictionary or a hash map in other languages.
// - All keys must be of the same type.
// - All values must be of the same type.
// - Keys are indexed, and you can iterate over them.
// - Use when you need to look up data by a key.
// - The set of keys can be dynamically changed (added or removed) at runtime.

// Use a struct when you want to group together different data types into a single, structured entity.
// - The fields (keys) are defined ahead of time at compile time.
// - Each field can have a different type.
// - You access fields with a dot notation (e.g., myStruct.myField).
// - Use when you have a fixed set of properties for an object or entity.

func main() {
	// More than one way of declaring a map in go
	// Maps == Dict ( PY )
	// All keys must be of the same type and all values as well , but keys and values can be of different type
	// map[string]string{} -> means that all keys are of type string and all values are of type string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Second way of declaring a map
	//var colors map[string]string

	// Another way
	// colors := make(map[string]string)

	// Adding elements to a map
	// colors["white"] = "#ffffff"
	// colors["red"] = "#ff0000"

	// calling an key from a map uses the []
	// colors["white"]

	// deleting key and values from a map - using the delete function.
	// delete(colors, "red")

	// fmt.Println(colors)
	printMap(colors)
}

// Iterate over all key, value pairs
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
