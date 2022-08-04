package main

import "fmt"

func main() {
	fmt.Println("----------Maps----------")

	// var colors map[string]string
	// colors["green"] = "#jh6564"
	// colors["white"] = "#000000"

	// colors := make(map[string]string)
	// colors["red"] = "#ff0000"
	// colors["blue"] = "#df4545"
	// delete(colors, "red")
		
	colors := map[string]string{
		"red":    "#ff0000",
		"blue":   "#dgf555",
		"yellow": "#hf5353",
	}
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is: ", hex)
	}
}
