package main

import "fmt"

func main() {
	fmt.Println("Five of Spades")
	// var cards string = "Five of Spades"
	// cards := "Five of Spades"

	// cards := deck{"Five of hearts", "Three Diamonds", newCard()}   //type deck is of []string from deck.go
	// cards = append(cards, "Six of Hearts")
	// fmt.Println(cards)
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// // cards.print()

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))

	fmt.Println(cards.toString())
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

// func newCard() string {
// 	return "Six of Spades"
// }
