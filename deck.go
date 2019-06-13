package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*Create a new type of deck
which is a slice of strings*/
type deck []card

//type deck methods -- methods called from a deck variable using the '.' key
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card.value, "of", card.suit)
	}
}

func (d deck) toString() string {
	var stringSliceOfDeck []string
	for _, c := range d {
		stringSliceOfDeck = append(stringSliceOfDeck, c.value+" of "+c.suit)
	}
	return strings.Join([]string(stringSliceOfDeck), ", ")
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//other deck functions, can be called without an existing deck variable
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{suit: suit, value: value})
		}
	}
	return cards
}

func newDeckFromFile(fileName string) deck {
	cards := deck{}
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := []string(strings.Split(string(bs), ","))
	for _, s := range ss {
		brokenDown := strings.Split(s, " ")
		if brokenDown[0] != "" {
			cards = append(cards, card{value: brokenDown[0], suit: brokenDown[2]})
		} else {
			cards = append(cards, card{value: brokenDown[1], suit: brokenDown[3]})
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
