package main

func main() {
	cards := newDeck()
	var hand deck
	cards.print()
	cards.shuffleDeck()
	cards.print()
	cards, hand = deal(cards, 5)
	hand.print()
	cards.print()
}
