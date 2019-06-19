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

//to run this project, type "go run main.go deck.go card.go" in the terminal.
//main.go is going to be changed to run the game on a listening http server in the near future
