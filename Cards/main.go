package main

import "fmt"

func main() {
	cards := newDeck()

	// hand, reaminingCards := deal(cards, 5)
	// hand.print()
	// reaminingCards.print()

	fmt.Println(cards.toString())

}

func newCard() string {
	return "Five of Hearts"
}
