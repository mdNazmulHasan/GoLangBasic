package main

func main() {
	cards := newDeck()
	cards = append(cards, "six of spades")
	cards.print()

}

func newCard() string {
	return "Five of diamonds"
}
