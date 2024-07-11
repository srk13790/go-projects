package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#ff00ff",
	// } //.....1

	// colors := make(map[int]string)

	// colors[0] = "White"
	// colors[1] = "Red"

	// delete(colors, 1) //.....2

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff00ff",
		"white": "#000000",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap (c map[string]string) {
	for key, val := range c {
		fmt.Println("Hex Code For " + val + " is " + key)
	}
}