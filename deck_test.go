package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(cards))
	}
	if cards[0].suit != "Hearts" && cards[0].value != "Ace" {
		t.Errorf("Expected first card of Ace of Hearts but got %v", cards[0])
	}
	if cards[len(cards)-1].suit != "Spades" && cards[len(cards)-1].value != "King" {
		t.Errorf("Expected last card of King of Spades but got %v", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck. Got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
