package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000FF",
		"white": "#ffffff",
	}

	// Add keys to maps
	// colors["white"] = "#ffffff"

	// Delete keys from maps
	// delete(colors, "blue")

	// Empty map
	// people := map[string]string

	// Alternative way to create empty map
	// testMap := make(map[string]string)

	// Iterate over maps
	printMap(colors)

	// fmt.Printf("%+v", colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
