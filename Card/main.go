package main

import "fmt"

func main() {
	fmt.Println("Five of Spades")
	// var cards string = "Five of Spades"
	cards := "Five of Spades"
	fmt.Println(cards)
	res := 0
	for i := 0; i < 10; i += 3 {
		res += i
	}
	fmt.Println(res)
}
