package main

func main() {
	cards := newDeckFromFile("my_cards")

	// hand, reaminingCards := deal(cards, 5)
	// hand.print()
	// reaminingCards.print()

	// cards.saveToFile("my_cards")

	cards.shuffle()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Hearts"
}
